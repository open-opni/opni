package bootstrap_test

import (
	"context"
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"net"
	"runtime"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	bootstrapv2 "github.com/open-panoptes/opni/pkg/apis/bootstrap/v2"
	"github.com/open-panoptes/opni/pkg/bootstrap"
	"github.com/open-panoptes/opni/pkg/ident"
	"github.com/open-panoptes/opni/pkg/storage"
	mock_ident "github.com/open-panoptes/opni/pkg/test/mock/ident"
	mock_storage "github.com/open-panoptes/opni/pkg/test/mock/storage"
	"github.com/open-panoptes/opni/pkg/test/testdata"
	"github.com/open-panoptes/opni/pkg/test/testutil"
	"github.com/open-panoptes/opni/pkg/tokens"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ = Describe("Client V2", Ordered, Label("unit"), func() {
	var fooIdent ident.Provider
	var cert *tls.Certificate
	var store storage.Backend
	var endpoint string

	BeforeAll(func() {
		if runtime.GOOS != "linux" {
			Skip("skipping test on non-linux OS")
		}
		fooIdent = mock_ident.NewTestIdentProvider(ctrl, "foo")
		var err error
		crt, err := tls.X509KeyPair(testdata.TestData("self_signed_leaf.crt"), testdata.TestData("self_signed_leaf.key"))
		Expect(err).NotTo(HaveOccurred())
		crt.Leaf, err = x509.ParseCertificate(crt.Certificate[0])
		Expect(err).NotTo(HaveOccurred())
		cert = &crt
		store = mock_storage.NewTestStorageBackend(context.Background(), ctrl)

		srv := grpc.NewServer(grpc.Creds(credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{*cert},
		})))

		server := bootstrap.NewServerV2(store, cert.PrivateKey.(crypto.Signer))
		bootstrapv2.RegisterBootstrapServer(srv, server)

		listener, err := net.Listen("tcp4", "127.0.0.1:0")
		Expect(err).NotTo(HaveOccurred())
		endpoint = listener.Addr().String()

		go srv.Serve(listener)

		DeferCleanup(func() {
			srv.Stop()
		})
	})

	It("should bootstrap with the server", func() {
		token, _ := store.CreateToken(context.Background(), 1*time.Minute)
		cc := bootstrap.ClientConfigV2{
			Token:         testutil.Must(tokens.FromBootstrapToken(token)),
			Endpoint:      endpoint,
			TrustStrategy: pkpTrustStrategy(cert.Leaf),
		}

		_, err := cc.Bootstrap(context.Background(), fooIdent)
		Expect(err).NotTo(HaveOccurred())
	})
	Context("error handling", func() {
		When("no token is given", func() {
			It("should error", func() {
				cc := bootstrap.ClientConfigV2{}
				kr, err := cc.Bootstrap(context.Background(), fooIdent)
				Expect(kr).To(BeNil())
				Expect(err).To(MatchError(bootstrap.ErrNoToken))
			})
		})
		When("an invalid endpoint is given", func() {
			It("should error", func() {
				token, _ := store.CreateToken(context.Background(), 1*time.Minute)
				cc := bootstrap.ClientConfigV2{
					Token:         testutil.Must(tokens.FromBootstrapToken(token)),
					Endpoint:      "\x7f",
					TrustStrategy: pkpTrustStrategy(cert.Leaf),
				}
				kr, err := cc.Bootstrap(context.Background(), fooIdent)
				Expect(kr).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("net/url"))
			})
		})
		When("the client fails to send a request to the server", func() {
			It("should error", func() {
				token, _ := store.CreateToken(context.Background(), 1*time.Minute)
				cc := bootstrap.ClientConfigV2{
					Token:         testutil.Must(tokens.FromBootstrapToken(token)),
					Endpoint:      "localhost:65545",
					TrustStrategy: pkpTrustStrategy(cert.Leaf),
				}
				kr, err := cc.Bootstrap(context.Background(), fooIdent)
				Expect(kr).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("invalid port"))
			})
		})
	})
})
