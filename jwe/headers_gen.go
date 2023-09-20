// Code generated by tools/cmd/genjwe/main.go. DO NOT EDIT.

package jwe

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/jwx/v3/cert"
	"github.com/lestrrat-go/jwx/v3/internal/base64"
	"github.com/lestrrat-go/jwx/v3/internal/json"
	"github.com/lestrrat-go/jwx/v3/internal/pool"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
)

const (
	AgreementPartyUInfoKey    = "apu"
	AgreementPartyVInfoKey    = "apv"
	AlgorithmKey              = "alg"
	CompressionKey            = "zip"
	ContentEncryptionKey      = "enc"
	ContentTypeKey            = "cty"
	CriticalKey               = "crit"
	EphemeralPublicKeyKey     = "epk"
	JWKKey                    = "jwk"
	JWKSetURLKey              = "jku"
	KeyIDKey                  = "kid"
	TypeKey                   = "typ"
	X509CertChainKey          = "x5c"
	X509CertThumbprintKey     = "x5t"
	X509CertThumbprintS256Key = "x5t#S256"
	X509URLKey                = "x5u"
)

// Headers describe a standard Header set.
type Headers interface {
	json.Marshaler
	json.Unmarshaler
	AgreementPartyUInfo() []byte
	AgreementPartyVInfo() []byte
	Algorithm() jwa.KeyEncryptionAlgorithm
	Compression() jwa.CompressionAlgorithm
	ContentEncryption() jwa.ContentEncryptionAlgorithm
	ContentType() string
	Critical() []string
	EphemeralPublicKey() jwk.Key
	JWK() jwk.Key
	JWKSetURL() string
	KeyID() string
	Type() string
	X509CertChain() *cert.Chain
	X509CertThumbprint() string
	X509CertThumbprintS256() string
	X509URL() string
	Iterate(ctx context.Context) Iterator
	Walk(ctx context.Context, v Visitor) error
	AsMap(ctx context.Context) (map[string]interface{}, error)
	Get(string) (interface{}, bool)
	Set(string, interface{}) error
	Remove(string) error
	Encode() ([]byte, error)
	Decode([]byte) error
	// PrivateParams returns the map containing the non-standard ('private') parameters
	// in the associated header. WARNING: DO NOT USE PrivateParams()
	// IF YOU HAVE CONCURRENT CODE ACCESSING THEM. Use AsMap() to
	// get a copy of the entire header instead
	PrivateParams() map[string]interface{}
	Clone(context.Context) (Headers, error)
	Copy(context.Context, Headers) error
	Merge(context.Context, Headers) (Headers, error)
}

type stdHeaders struct {
	agreementPartyUInfo    []byte
	agreementPartyVInfo    []byte
	algorithm              *jwa.KeyEncryptionAlgorithm
	compression            *jwa.CompressionAlgorithm
	contentEncryption      *jwa.ContentEncryptionAlgorithm
	contentType            *string
	critical               []string
	ephemeralPublicKey     jwk.Key
	jwk                    jwk.Key
	jwkSetURL              *string
	keyID                  *string
	typ                    *string
	x509CertChain          *cert.Chain
	x509CertThumbprint     *string
	x509CertThumbprintS256 *string
	x509URL                *string
	privateParams          map[string]interface{}
	mu                     *sync.RWMutex
}

func NewHeaders() Headers {
	return &stdHeaders{
		mu:            &sync.RWMutex{},
		privateParams: map[string]interface{}{},
	}
}

func (h *stdHeaders) AgreementPartyUInfo() []byte {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.agreementPartyUInfo
}

func (h *stdHeaders) AgreementPartyVInfo() []byte {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.agreementPartyVInfo
}

func (h *stdHeaders) Algorithm() jwa.KeyEncryptionAlgorithm {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.algorithm == nil {
		return ""
	}
	return *(h.algorithm)
}

func (h *stdHeaders) Compression() jwa.CompressionAlgorithm {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.compression == nil {
		return jwa.NoCompress
	}
	return *(h.compression)
}

func (h *stdHeaders) ContentEncryption() jwa.ContentEncryptionAlgorithm {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.contentEncryption == nil {
		return ""
	}
	return *(h.contentEncryption)
}

func (h *stdHeaders) ContentType() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.contentType == nil {
		return ""
	}
	return *(h.contentType)
}

func (h *stdHeaders) Critical() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.critical
}

func (h *stdHeaders) EphemeralPublicKey() jwk.Key {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.ephemeralPublicKey
}

func (h *stdHeaders) JWK() jwk.Key {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.jwk
}

func (h *stdHeaders) JWKSetURL() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.jwkSetURL == nil {
		return ""
	}
	return *(h.jwkSetURL)
}

func (h *stdHeaders) KeyID() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.keyID == nil {
		return ""
	}
	return *(h.keyID)
}

func (h *stdHeaders) Type() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.typ == nil {
		return ""
	}
	return *(h.typ)
}

func (h *stdHeaders) X509CertChain() *cert.Chain {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.x509CertChain
}

func (h *stdHeaders) X509CertThumbprint() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.x509CertThumbprint == nil {
		return ""
	}
	return *(h.x509CertThumbprint)
}

func (h *stdHeaders) X509CertThumbprintS256() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.x509CertThumbprintS256 == nil {
		return ""
	}
	return *(h.x509CertThumbprintS256)
}

func (h *stdHeaders) X509URL() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.x509URL == nil {
		return ""
	}
	return *(h.x509URL)
}

func (h *stdHeaders) makePairs() []*HeaderPair {
	h.mu.RLock()
	defer h.mu.RUnlock()
	var pairs []*HeaderPair
	if h.agreementPartyUInfo != nil {
		pairs = append(pairs, &HeaderPair{Key: AgreementPartyUInfoKey, Value: h.agreementPartyUInfo})
	}
	if h.agreementPartyVInfo != nil {
		pairs = append(pairs, &HeaderPair{Key: AgreementPartyVInfoKey, Value: h.agreementPartyVInfo})
	}
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
	}
	if h.compression != nil {
		pairs = append(pairs, &HeaderPair{Key: CompressionKey, Value: *(h.compression)})
	}
	if h.contentEncryption != nil {
		pairs = append(pairs, &HeaderPair{Key: ContentEncryptionKey, Value: *(h.contentEncryption)})
	}
	if h.contentType != nil {
		pairs = append(pairs, &HeaderPair{Key: ContentTypeKey, Value: *(h.contentType)})
	}
	if h.critical != nil {
		pairs = append(pairs, &HeaderPair{Key: CriticalKey, Value: h.critical})
	}
	if h.ephemeralPublicKey != nil {
		pairs = append(pairs, &HeaderPair{Key: EphemeralPublicKeyKey, Value: h.ephemeralPublicKey})
	}
	if h.jwk != nil {
		pairs = append(pairs, &HeaderPair{Key: JWKKey, Value: h.jwk})
	}
	if h.jwkSetURL != nil {
		pairs = append(pairs, &HeaderPair{Key: JWKSetURLKey, Value: *(h.jwkSetURL)})
	}
	if h.keyID != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyIDKey, Value: *(h.keyID)})
	}
	if h.typ != nil {
		pairs = append(pairs, &HeaderPair{Key: TypeKey, Value: *(h.typ)})
	}
	if h.x509CertChain != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertChainKey, Value: h.x509CertChain})
	}
	if h.x509CertThumbprint != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintKey, Value: *(h.x509CertThumbprint)})
	}
	if h.x509CertThumbprintS256 != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintS256Key, Value: *(h.x509CertThumbprintS256)})
	}
	if h.x509URL != nil {
		pairs = append(pairs, &HeaderPair{Key: X509URLKey, Value: *(h.x509URL)})
	}
	for k, v := range h.privateParams {
		pairs = append(pairs, &HeaderPair{Key: k, Value: v})
	}
	return pairs
}

func (h *stdHeaders) PrivateParams() map[string]interface{} {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.privateParams
}

func (h *stdHeaders) Get(name string) (interface{}, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	switch name {
	case AgreementPartyUInfoKey:
		if h.agreementPartyUInfo == nil {
			return nil, false
		}
		return h.agreementPartyUInfo, true
	case AgreementPartyVInfoKey:
		if h.agreementPartyVInfo == nil {
			return nil, false
		}
		return h.agreementPartyVInfo, true
	case AlgorithmKey:
		if h.algorithm == nil {
			return nil, false
		}
		return *(h.algorithm), true
	case CompressionKey:
		if h.compression == nil {
			return nil, false
		}
		return *(h.compression), true
	case ContentEncryptionKey:
		if h.contentEncryption == nil {
			return nil, false
		}
		return *(h.contentEncryption), true
	case ContentTypeKey:
		if h.contentType == nil {
			return nil, false
		}
		return *(h.contentType), true
	case CriticalKey:
		if h.critical == nil {
			return nil, false
		}
		return h.critical, true
	case EphemeralPublicKeyKey:
		if h.ephemeralPublicKey == nil {
			return nil, false
		}
		return h.ephemeralPublicKey, true
	case JWKKey:
		if h.jwk == nil {
			return nil, false
		}
		return h.jwk, true
	case JWKSetURLKey:
		if h.jwkSetURL == nil {
			return nil, false
		}
		return *(h.jwkSetURL), true
	case KeyIDKey:
		if h.keyID == nil {
			return nil, false
		}
		return *(h.keyID), true
	case TypeKey:
		if h.typ == nil {
			return nil, false
		}
		return *(h.typ), true
	case X509CertChainKey:
		if h.x509CertChain == nil {
			return nil, false
		}
		return h.x509CertChain, true
	case X509CertThumbprintKey:
		if h.x509CertThumbprint == nil {
			return nil, false
		}
		return *(h.x509CertThumbprint), true
	case X509CertThumbprintS256Key:
		if h.x509CertThumbprintS256 == nil {
			return nil, false
		}
		return *(h.x509CertThumbprintS256), true
	case X509URLKey:
		if h.x509URL == nil {
			return nil, false
		}
		return *(h.x509URL), true
	default:
		v, ok := h.privateParams[name]
		return v, ok
	}
}

func (h *stdHeaders) Set(name string, value interface{}) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.setNoLock(name, value)
}

func (h *stdHeaders) setNoLock(name string, value interface{}) error {
	switch name {
	case AgreementPartyUInfoKey:
		if v, ok := value.([]byte); ok {
			h.agreementPartyUInfo = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, AgreementPartyUInfoKey, value)
	case AgreementPartyVInfoKey:
		if v, ok := value.([]byte); ok {
			h.agreementPartyVInfo = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, AgreementPartyVInfoKey, value)
	case AlgorithmKey:
		if v, ok := value.(jwa.KeyEncryptionAlgorithm); ok {
			h.algorithm = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, AlgorithmKey, value)
	case CompressionKey:
		if v, ok := value.(jwa.CompressionAlgorithm); ok {
			h.compression = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, CompressionKey, value)
	case ContentEncryptionKey:
		if v, ok := value.(jwa.ContentEncryptionAlgorithm); ok {
			if v == "" {
				return fmt.Errorf(`"enc" field cannot be an empty string`)
			}
			h.contentEncryption = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, ContentEncryptionKey, value)
	case ContentTypeKey:
		if v, ok := value.(string); ok {
			h.contentType = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, ContentTypeKey, value)
	case CriticalKey:
		if v, ok := value.([]string); ok {
			h.critical = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, CriticalKey, value)
	case EphemeralPublicKeyKey:
		if v, ok := value.(jwk.Key); ok {
			h.ephemeralPublicKey = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, EphemeralPublicKeyKey, value)
	case JWKKey:
		if v, ok := value.(jwk.Key); ok {
			h.jwk = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, JWKKey, value)
	case JWKSetURLKey:
		if v, ok := value.(string); ok {
			h.jwkSetURL = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, JWKSetURLKey, value)
	case KeyIDKey:
		if v, ok := value.(string); ok {
			h.keyID = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, KeyIDKey, value)
	case TypeKey:
		if v, ok := value.(string); ok {
			h.typ = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, TypeKey, value)
	case X509CertChainKey:
		if v, ok := value.(*cert.Chain); ok {
			h.x509CertChain = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertChainKey, value)
	case X509CertThumbprintKey:
		if v, ok := value.(string); ok {
			h.x509CertThumbprint = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertThumbprintKey, value)
	case X509CertThumbprintS256Key:
		if v, ok := value.(string); ok {
			h.x509CertThumbprintS256 = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertThumbprintS256Key, value)
	case X509URLKey:
		if v, ok := value.(string); ok {
			h.x509URL = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509URLKey, value)
	default:
		if h.privateParams == nil {
			h.privateParams = map[string]interface{}{}
		}
		h.privateParams[name] = value
	}
	return nil
}

func (h *stdHeaders) Remove(key string) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	switch key {
	case AgreementPartyUInfoKey:
		h.agreementPartyUInfo = nil
	case AgreementPartyVInfoKey:
		h.agreementPartyVInfo = nil
	case AlgorithmKey:
		h.algorithm = nil
	case CompressionKey:
		h.compression = nil
	case ContentEncryptionKey:
		h.contentEncryption = nil
	case ContentTypeKey:
		h.contentType = nil
	case CriticalKey:
		h.critical = nil
	case EphemeralPublicKeyKey:
		h.ephemeralPublicKey = nil
	case JWKKey:
		h.jwk = nil
	case JWKSetURLKey:
		h.jwkSetURL = nil
	case KeyIDKey:
		h.keyID = nil
	case TypeKey:
		h.typ = nil
	case X509CertChainKey:
		h.x509CertChain = nil
	case X509CertThumbprintKey:
		h.x509CertThumbprint = nil
	case X509CertThumbprintS256Key:
		h.x509CertThumbprintS256 = nil
	case X509URLKey:
		h.x509URL = nil
	default:
		delete(h.privateParams, key)
	}
	return nil
}

func (h *stdHeaders) UnmarshalJSON(buf []byte) error {
	h.agreementPartyUInfo = nil
	h.agreementPartyVInfo = nil
	h.algorithm = nil
	h.compression = nil
	h.contentEncryption = nil
	h.contentType = nil
	h.critical = nil
	h.ephemeralPublicKey = nil
	h.jwk = nil
	h.jwkSetURL = nil
	h.keyID = nil
	h.typ = nil
	h.x509CertChain = nil
	h.x509CertThumbprint = nil
	h.x509CertThumbprintS256 = nil
	h.x509URL = nil
	dec := json.NewDecoder(bytes.NewReader(buf))
LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf(`error reading token: %w`, err)
		}
		switch tok := tok.(type) {
		case json.Delim:
			// Assuming we're doing everything correctly, we should ONLY
			// get either '{' or '}' here.
			if tok == '}' { // End of object
				break LOOP
			} else if tok != '{' {
				return fmt.Errorf(`expected '{', but got '%c'`, tok)
			}
		case string: // Objects can only have string keys
			switch tok {
			case AgreementPartyUInfoKey:
				if err := json.AssignNextBytesToken(&h.agreementPartyUInfo, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, AgreementPartyUInfoKey, err)
				}
			case AgreementPartyVInfoKey:
				if err := json.AssignNextBytesToken(&h.agreementPartyVInfo, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, AgreementPartyVInfoKey, err)
				}
			case AlgorithmKey:
				var decoded jwa.KeyEncryptionAlgorithm
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, AlgorithmKey, err)
				}
				h.algorithm = &decoded
			case CompressionKey:
				var decoded jwa.CompressionAlgorithm
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, CompressionKey, err)
				}
				h.compression = &decoded
			case ContentEncryptionKey:
				var decoded jwa.ContentEncryptionAlgorithm
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, ContentEncryptionKey, err)
				}
				h.contentEncryption = &decoded
			case ContentTypeKey:
				if err := json.AssignNextStringToken(&h.contentType, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, ContentTypeKey, err)
				}
			case CriticalKey:
				var decoded []string
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, CriticalKey, err)
				}
				h.critical = decoded
			case EphemeralPublicKeyKey:
				var buf json.RawMessage
				if err := dec.Decode(&buf); err != nil {
					return fmt.Errorf(`failed to decode value for key %s:%w`, EphemeralPublicKeyKey, err)
				}
				key, err := jwk.ParseKey(buf)
				if err != nil {
					return fmt.Errorf(`failed to parse JWK for key %s: %w`, EphemeralPublicKeyKey, err)
				}
				h.ephemeralPublicKey = key
			case JWKKey:
				var buf json.RawMessage
				if err := dec.Decode(&buf); err != nil {
					return fmt.Errorf(`failed to decode value for key %s:%w`, JWKKey, err)
				}
				key, err := jwk.ParseKey(buf)
				if err != nil {
					return fmt.Errorf(`failed to parse JWK for key %s: %w`, JWKKey, err)
				}
				h.jwk = key
			case JWKSetURLKey:
				if err := json.AssignNextStringToken(&h.jwkSetURL, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, JWKSetURLKey, err)
				}
			case KeyIDKey:
				if err := json.AssignNextStringToken(&h.keyID, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, KeyIDKey, err)
				}
			case TypeKey:
				if err := json.AssignNextStringToken(&h.typ, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, TypeKey, err)
				}
			case X509CertChainKey:
				var decoded cert.Chain
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertChainKey, err)
				}
				h.x509CertChain = &decoded
			case X509CertThumbprintKey:
				if err := json.AssignNextStringToken(&h.x509CertThumbprint, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertThumbprintKey, err)
				}
			case X509CertThumbprintS256Key:
				if err := json.AssignNextStringToken(&h.x509CertThumbprintS256, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertThumbprintS256Key, err)
				}
			case X509URLKey:
				if err := json.AssignNextStringToken(&h.x509URL, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509URLKey, err)
				}
			default:
				decoded, err := registry.Decode(dec, tok)
				if err != nil {
					return err
				}
				h.setNoLock(tok, decoded)
			}
		default:
			return fmt.Errorf(`invalid token %T`, tok)
		}
	}
	return nil
}

func (h stdHeaders) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	fields := make([]string, 0, 16)
	for _, pair := range h.makePairs() {
		fields = append(fields, pair.Key.(string))
		data[pair.Key.(string)] = pair.Value
	}

	sort.Strings(fields)
	buf := pool.GetBytesBuffer()
	defer pool.ReleaseBytesBuffer(buf)
	buf.WriteByte('{')
	enc := json.NewEncoder(buf)
	for i, f := range fields {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune('"')
		buf.WriteString(f)
		buf.WriteString(`":`)
		v := data[f]
		switch v := v.(type) {
		case []byte:
			buf.WriteRune('"')
			buf.WriteString(base64.EncodeToString(v))
			buf.WriteRune('"')
		default:
			if err := enc.Encode(v); err != nil {
				return nil, fmt.Errorf(`failed to encode value for field %s`, f)
			}
			buf.Truncate(buf.Len() - 1)
		}
	}
	buf.WriteByte('}')
	ret := make([]byte, buf.Len())
	copy(ret, buf.Bytes())
	return ret, nil
}
