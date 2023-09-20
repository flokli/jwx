// Code generated by tools/cmd/genoptions/main.go. DO NOT EDIT.

package jwt

import (
	"context"
	"io/fs"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwe"
	"github.com/lestrrat-go/jwx/v3/jws"
	"github.com/lestrrat-go/option"
)

type Option = option.Interface

// EncryptOption describes an Option that can be passed to (jwt.Serializer).Encrypt
type EncryptOption interface {
	Option
	encryptOption()
}

type encryptOption struct {
	Option
}

func (*encryptOption) encryptOption() {}

// GlobalOption describes an Option that can be passed to `Settings()`.
type GlobalOption interface {
	Option
	globalOption()
}

type globalOption struct {
	Option
}

func (*globalOption) globalOption() {}

// ParseOption describes an Option that can be passed to `jwt.Parse()`.
// ParseOption also implements ReadFileOption, therefore it may be
// safely pass them to `jwt.ReadFile()`
type ParseOption interface {
	Option
	parseOption()
	readFileOption()
}

type parseOption struct {
	Option
}

func (*parseOption) parseOption() {}

func (*parseOption) readFileOption() {}

// ReadFileOption is a type of `Option` that can be passed to `jws.ReadFile`
type ReadFileOption interface {
	Option
	readFileOption()
}

type readFileOption struct {
	Option
}

func (*readFileOption) readFileOption() {}

// SignEncryptParseOption describes an Option that can be passed to both `jwt.Sign()` or
// `jwt.Parse()`
type SignEncryptParseOption interface {
	Option
	parseOption()
	encryptOption()
	readFileOption()
	signOption()
}

type signEncryptParseOption struct {
	Option
}

func (*signEncryptParseOption) parseOption() {}

func (*signEncryptParseOption) encryptOption() {}

func (*signEncryptParseOption) readFileOption() {}

func (*signEncryptParseOption) signOption() {}

// SignOption describes an Option that can be passed to `jwt.Sign()` or
// (jwt.Serializer).Sign
type SignOption interface {
	Option
	signOption()
}

type signOption struct {
	Option
}

func (*signOption) signOption() {}

// ValidateOption describes an Option that can be passed to Validate().
// ValidateOption also implements ParseOption, therefore it may be
// safely passed to `Parse()` (and thus `jwt.ReadFile()`)
type ValidateOption interface {
	Option
	parseOption()
	readFileOption()
	validateOption()
}

type validateOption struct {
	Option
}

func (*validateOption) parseOption() {}

func (*validateOption) readFileOption() {}

func (*validateOption) validateOption() {}

type identAcceptableSkew struct{}
type identClock struct{}
type identContext struct{}
type identEncryptOption struct{}
type identFS struct{}
type identFlattenAudience struct{}
type identFormKey struct{}
type identHeaderKey struct{}
type identKeyProvider struct{}
type identNumericDateFormatPrecision struct{}
type identNumericDateParsePedantic struct{}
type identNumericDateParsePrecision struct{}
type identPedantic struct{}
type identSignOption struct{}
type identToken struct{}
type identTruncation struct{}
type identValidate struct{}
type identValidator struct{}
type identVerify struct{}

func (identAcceptableSkew) String() string {
	return "WithAcceptableSkew"
}

func (identClock) String() string {
	return "WithClock"
}

func (identContext) String() string {
	return "WithContext"
}

func (identEncryptOption) String() string {
	return "WithEncryptOption"
}

func (identFS) String() string {
	return "WithFS"
}

func (identFlattenAudience) String() string {
	return "WithFlattenAudience"
}

func (identFormKey) String() string {
	return "WithFormKey"
}

func (identHeaderKey) String() string {
	return "WithHeaderKey"
}

func (identKeyProvider) String() string {
	return "WithKeyProvider"
}

func (identNumericDateFormatPrecision) String() string {
	return "WithNumericDateFormatPrecision"
}

func (identNumericDateParsePedantic) String() string {
	return "WithNumericDateParsePedantic"
}

func (identNumericDateParsePrecision) String() string {
	return "WithNumericDateParsePrecision"
}

func (identPedantic) String() string {
	return "WithPedantic"
}

func (identSignOption) String() string {
	return "WithSignOption"
}

func (identToken) String() string {
	return "WithToken"
}

func (identTruncation) String() string {
	return "WithTruncation"
}

func (identValidate) String() string {
	return "WithValidate"
}

func (identValidator) String() string {
	return "WithValidator"
}

func (identVerify) String() string {
	return "WithVerify"
}

// WithAcceptableSkew specifies the duration in which exp and nbf
// claims may differ by. This value should be positive
func WithAcceptableSkew(v time.Duration) ValidateOption {
	return &validateOption{option.New(identAcceptableSkew{}, v)}
}

// WithClock specifies the `Clock` to be used when verifying
// exp and nbf claims.
func WithClock(v Clock) ValidateOption {
	return &validateOption{option.New(identClock{}, v)}
}

// WithContext allows you to specify a context.Context object to be used
// with `jwt.Validate()` option.
//
// Please be aware that in the next major release of this library,
// `jwt.Validate()`'s signature will change to include an explicit
// `context.Context` object.
func WithContext(v context.Context) ValidateOption {
	return &validateOption{option.New(identContext{}, v)}
}

// WithEncryptOption provides an escape hatch for cases where extra options to
// `(jws.Serializer).Encrypt()` must be specified when usng `jwt.Sign()`. Normally you do not
// need to use this.
func WithEncryptOption(v jwe.EncryptOption) EncryptOption {
	return &encryptOption{option.New(identEncryptOption{}, v)}
}

// WithFS specifies the source `fs.FS` object to read the file from.
func WithFS(v fs.FS) ReadFileOption {
	return &readFileOption{option.New(identFS{}, v)}
}

// WithFlattenAudience specifies the the `jwt.FlattenAudience` option on
// every token defaults to enabled. You can still disable this on a per-object
// basis using the `jwt.Options().Disable(jwt.FlattenAudience)` method call.
//
// See the documentation for `jwt.TokenOptionSet`, `(jwt.Token).Options`, and
// `jwt.FlattenAudience` for more details
func WithFlattenAudience(v bool) GlobalOption {
	return &globalOption{option.New(identFlattenAudience{}, v)}
}

// WithFormKey is used to specify header keys to search for tokens.
//
// While the type system allows this option to be passed to jwt.Parse() directly,
// doing so will have no effect. Only use it for HTTP request parsing functions
func WithFormKey(v string) ParseOption {
	return &parseOption{option.New(identFormKey{}, v)}
}

// WithHeaderKey is used to specify header keys to search for tokens.
//
// While the type system allows this option to be passed to `jwt.Parse()` directly,
// doing so will have no effect. Only use it for HTTP request parsing functions
func WithHeaderKey(v string) ParseOption {
	return &parseOption{option.New(identHeaderKey{}, v)}
}

// WithKeyProvider allows users to specify an object to provide keys to
// sign/verify tokens using arbitrary code. Please read the documentation
// for `jws.KeyProvider` in the `jws` package for details on how this works.
func WithKeyProvider(v jws.KeyProvider) ParseOption {
	return &parseOption{option.New(identKeyProvider{}, v)}
}

// WithNumericDateFormatPrecision sets the precision up to which the
// library uses to format fractional dates found in the numeric date
// fields. Default is 0 (second, no fractionals), max is 9 (nanosecond)
func WithNumericDateFormatPrecision(v int) GlobalOption {
	return &globalOption{option.New(identNumericDateFormatPrecision{}, v)}
}

// WithNumericDateParsePedantic specifies if the parser should behave
// in a pedantic manner when parsing numeric dates. Normally this library
// attempts to interpret timestamps as a numeric value representing
// number of seconds (with an optional fractional part), but if that fails
// it tries to parse using a RFC3339 parser. This allows us to parse
// payloads from non-comforming servers.
//
// However, when you set WithNumericDateParePedantic to `true`, the
// RFC3339 parser is not tried, and we expect a numeric value strictly
func WithNumericDateParsePedantic(v bool) GlobalOption {
	return &globalOption{option.New(identNumericDateParsePedantic{}, v)}
}

// WithNumericDateParsePrecision sets the precision up to which the
// library uses to parse fractional dates found in the numeric date
// fields. Default is 0 (second, no fractionals), max is 9 (nanosecond)
func WithNumericDateParsePrecision(v int) GlobalOption {
	return &globalOption{option.New(identNumericDateParsePrecision{}, v)}
}

// WithPedantic enables pedantic mode for parsing JWTs. Currently this only
// applies to checking for the correct `typ` and/or `cty` when necessary.
func WithPedantic(v bool) ParseOption {
	return &parseOption{option.New(identPedantic{}, v)}
}

// WithSignOption provides an escape hatch for cases where extra options to
// `jws.Sign()` must be specified when usng `jwt.Sign()`. Normally you do not
// need to use this.
func WithSignOption(v jws.SignOption) SignOption {
	return &signOption{option.New(identSignOption{}, v)}
}

// WithToken specifies the token instance where the result JWT is stored
// when parsing JWT tokensthat is used when parsing
func WithToken(v Token) ParseOption {
	return &parseOption{option.New(identToken{}, v)}
}

// WithTruncation speficies the amount that should be used when
// truncating time values used during time-based validation routines.
// By default time values are truncated down to second accuracy.
// If you want to use sub-second accuracy, you will need to set
// this value to 0.
func WithTruncation(v time.Duration) ValidateOption {
	return &validateOption{option.New(identTruncation{}, v)}
}

// WithValidate is passed to `Parse()` method to denote that the
// validation of the JWT token should be performed (or not) after
// a successful parsing of the incoming payload.
//
// This option is enabled by default.
//
// If you would like disable validation,
// you must use `jwt.WithValidate(false)` or use `jwt.ParseInsecure()`
func WithValidate(v bool) ParseOption {
	return &parseOption{option.New(identValidate{}, v)}
}

// WithValidator validates the token with the given Validator.
//
// For example, in order to validate tokens that are only valid during August, you would write
//
//	validator := jwt.ValidatorFunc(func(_ context.Context, t jwt.Token) error {
//		if time.Now().Month() != 8 {
//			return fmt.Errorf(`tokens are only valid during August!`)
//		}
//		return nil
//	})
//	err := jwt.Validate(token, jwt.WithValidator(validator))
func WithValidator(v Validator) ValidateOption {
	return &validateOption{option.New(identValidator{}, v)}
}

// WithVerify is passed to `Parse()` method to denote that the
// signature verification should be performed after a successful
// deserialization of the incoming payload.
//
// This option is enabled by default.
//
// If you do not provide any verification key sources, `jwt.Parse()`
// would return an error.
//
// If you would like to only parse the JWT payload and not verify it,
// you must use `jwt.WithVerify(false)` or use `jwt.ParseInsecure()`
func WithVerify(v bool) ParseOption {
	return &parseOption{option.New(identVerify{}, v)}
}
