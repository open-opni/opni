package tokens

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
)

var ErrMalformedToken = errors.New("malformed token")

type Token struct {
	ID     []byte `json:"id"`               // bytes 0-5
	Secret []byte `json:"secret,omitempty"` // bytes 6-31
}

// Creates a new token by reading bytes from the given random source.
// the default source is crypto/rand.Reader.
func NewToken(source ...io.Reader) *Token {
	entropy := rand.Reader
	if len(source) > 0 {
		entropy = source[0]
	}
	buf := make([]byte, 32)
	if _, err := io.ReadFull(entropy, buf); err != nil {
		panic(err)
	}
	return &Token{
		ID:     buf[:6],
		Secret: buf[6:],
	}
}

func (t *Token) EncodeJSON() []byte {
	data, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return data
}

func (t *Token) EncodeHex() string {
	return hex.EncodeToString(t.ID) + "." + hex.EncodeToString(t.Secret)
}

func (t *Token) HexID() string {
	return hex.EncodeToString(t.ID)
}

func (t *Token) HexSecret() string {
	return hex.EncodeToString(t.Secret)
}

func (t *Token) Reference() *corev1.Reference {
	return &corev1.Reference{
		Id: t.HexID(),
	}
}

func ParseJSON(data []byte) (*Token, error) {
	t := &Token{}
	if err := json.Unmarshal(data, t); err != nil {
		return nil, err
	}
	return t, nil
}

func ParseHex(str string) (*Token, error) {
	parts := bytes.Split([]byte(str), []byte("."))
	if len(parts) != 2 ||
		len(parts[0]) != hex.EncodedLen(6) ||
		len(parts[1]) != hex.EncodedLen(26) {
		return nil, ErrMalformedToken
	}
	t := &Token{
		ID:     make([]byte, 6),
		Secret: make([]byte, 26),
	}
	if n, err := hex.Decode(t.ID, parts[0]); err != nil || n != 6 {
		return nil, ErrMalformedToken
	}
	if n, err := hex.Decode(t.Secret, parts[1]); err != nil || n != 26 {
		return nil, ErrMalformedToken
	}
	return t, nil
}

// Signs the token and returns a JWS with the payload detached
func (t *Token) SignDetached(key interface{}) ([]byte, error) {
	if _, ok := key.(ed25519.PrivateKey); !ok {
		return nil, errors.New("invalid key type, expected ed25519.PrivateKey")
	}
	jsonData, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	sig, err := jws.Sign(jsonData, jwa.EdDSA, key)
	if err != nil {
		return nil, err
	}
	firstIndex := bytes.IndexByte(sig, '.')
	lastIndex := bytes.LastIndexByte(sig, '.')
	buf := new(bytes.Buffer)
	buf.Write(sig[:firstIndex+1])
	buf.Write(sig[lastIndex:])
	return buf.Bytes(), nil
}

// Verifies a JWS with a detached signature. If the signature is valid,
// also returns the complete message with re-attached payload.
func (t *Token) VerifyDetached(sig []byte, key interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	firstIndex := bytes.IndexByte(sig, '.')
	lastIndex := bytes.LastIndexByte(sig, '.')
	if firstIndex == -1 || lastIndex == -1 {
		return nil, ErrMalformedToken
	}
	payload := base64.RawURLEncoding.EncodeToString(jsonData)
	buf := new(bytes.Buffer)
	buf.Write(sig[:firstIndex+1])
	buf.WriteString(payload)
	buf.Write(sig[lastIndex:])
	fullToken := buf.Bytes()
	cloned := make([]byte, len(fullToken))
	copy(cloned, fullToken)
	_, err = jws.Verify(cloned, jwa.EdDSA, key)
	if err != nil {
		return nil, err
	}
	return fullToken, nil
}
