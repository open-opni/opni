package bootstrap

import (
	"context"
	"crypto/x509"
	"fmt"

	bootstrapv1 "github.com/open-panoptes/opni/pkg/apis/bootstrap/v1"
	"github.com/open-panoptes/opni/pkg/auth"
	"github.com/open-panoptes/opni/pkg/ecdh"
	"github.com/open-panoptes/opni/pkg/ident"
	"github.com/open-panoptes/opni/pkg/keyring"
	"github.com/open-panoptes/opni/pkg/tokens"
	"github.com/open-panoptes/opni/pkg/trust"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type ClientConfig struct {
	Capability    string
	Token         *tokens.Token
	Endpoint      string
	DialOpts      []grpc.DialOption
	TrustStrategy trust.Strategy
}

func (c *ClientConfig) Bootstrap(
	ctx context.Context,
	ident ident.Provider,
) (keyring.Keyring, error) {
	if c.Token == nil {
		return nil, ErrNoToken
	}
	response, serverLeafCert, err := c.bootstrapJoin(ctx)
	if err != nil {
		return nil, err
	}

	completeJws, err := c.findValidSignature(
		response.Signatures, serverLeafCert.PublicKey)
	if err != nil {
		return nil, err
	}

	tlsConfig, err := c.TrustStrategy.TLSConfig()
	if err != nil {
		return nil, err
	}

	cc, err := grpc.DialContext(ctx, c.Endpoint,
		append(c.DialOpts,
			grpc.WithBlock(),
			grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
			grpc.WithChainStreamInterceptor(otelgrpc.StreamClientInterceptor()),
			grpc.WithChainUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		)...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial gateway: %w", err)
	}
	defer cc.Close()

	client := bootstrapv1.NewBootstrapClient(cc)

	ekp := ecdh.NewEphemeralKeyPair()
	id, err := ident.UniqueIdentifier(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain unique identifier: %w", err)
	}
	authReq := &bootstrapv1.BootstrapAuthRequest{
		ClientID:     id,
		ClientPubKey: ekp.PublicKey.Bytes(),
		Capability:   c.Capability,
	}

	authResp, err := client.Auth(metadata.NewOutgoingContext(ctx, metadata.Pairs(
		auth.AuthorizationKey, "Bearer "+string(completeJws),
	)), authReq)
	if err != nil {
		return nil, fmt.Errorf("auth request failed: %w", err)
	}

	serverPubKey, err := ecdh.ServerPubKey(authResp)
	if err != nil {
		return nil, err
	}
	sharedSecret, err := ecdh.DeriveSharedSecret(ekp, serverPubKey)
	if err != nil {
		return nil, err
	}

	keys := []any{keyring.NewSharedKeys(sharedSecret)}
	if k := c.TrustStrategy.PersistentKey(); k != nil {
		keys = append(keys, k)
	}
	return keyring.New(keys...), nil
}

func (c *ClientConfig) bootstrapJoin(ctx context.Context) (*bootstrapv1.BootstrapJoinResponse, *x509.Certificate, error) {
	tlsConfig, err := c.TrustStrategy.TLSConfig()
	if err != nil {
		return nil, nil, err
	}
	cc, err := grpc.DialContext(ctx, c.Endpoint,
		append(c.DialOpts,
			grpc.WithBlock(),
			grpc.FailOnNonTempDialError(true),
			grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
			grpc.WithChainStreamInterceptor(otelgrpc.StreamClientInterceptor()),
			grpc.WithChainUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		)...,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial gateway: %w", err)
	}
	defer cc.Close()

	client := bootstrapv1.NewBootstrapClient(cc)

	var peer peer.Peer
	resp, err := client.Join(ctx, &bootstrapv1.BootstrapJoinRequest{}, grpc.Peer(&peer), grpc.WaitForReady(true))
	if err != nil {
		return nil, nil, fmt.Errorf("join request failed: %w", err)
	}

	tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo)
	if !ok {
		return nil, nil, fmt.Errorf("unexpected type for peer TLS info: %#T", peer.AuthInfo)
	}

	return resp, tlsInfo.State.PeerCertificates[0], nil
}

func (c *ClientConfig) findValidSignature(
	signatures map[string][]byte,
	pubKey interface{},
) ([]byte, error) {
	if sig, ok := signatures[c.Token.HexID()]; ok {
		return c.Token.VerifyDetached(sig, pubKey)
	}
	return nil, ErrNoValidSignature
}
