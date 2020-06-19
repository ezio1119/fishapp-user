// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _user_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Email

	// no validation rules for Name

	// no validation rules for Introduction

	// no validation rules for Sex

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on TokenPair with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TokenPair) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IdToken

	// no validation rules for RefreshToken

	return nil
}

// TokenPairValidationError is the validation error returned by
// TokenPair.Validate if the designated constraints aren't met.
type TokenPairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenPairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenPairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenPairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenPairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenPairValidationError) ErrorName() string { return "TokenPairValidationError" }

// Error satisfies the builtin error interface
func (e TokenPairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTokenPair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenPairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenPairValidationError{}

// Validate checks the field values on CurrentUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CurrentUserReq) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CurrentUserReqValidationError is the validation error returned by
// CurrentUserReq.Validate if the designated constraints aren't met.
type CurrentUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrentUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrentUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrentUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrentUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrentUserReqValidationError) ErrorName() string { return "CurrentUserReqValidationError" }

// Error satisfies the builtin error interface
func (e CurrentUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrentUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrentUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrentUserReqValidationError{}

// Validate checks the field values on GetUserReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetUserReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return GetUserReqValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// GetUserReqValidationError is the validation error returned by
// GetUserReq.Validate if the designated constraints aren't met.
type GetUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReqValidationError) ErrorName() string { return "GetUserReqValidationError" }

// Error satisfies the builtin error interface
func (e GetUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReqValidationError{}

// Validate checks the field values on CreateUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateUserReq) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Data.(type) {

	case *CreateUserReq_Info:

		if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateUserReqValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CreateUserReq_ImageChunk:

		if len(m.GetImageChunk()) > 65536 {
			return CreateUserReqValidationError{
				field:  "ImageChunk",
				reason: "value length must be at most 65536 bytes",
			}
		}

	default:
		return CreateUserReqValidationError{
			field:  "Data",
			reason: "value is required",
		}

	}

	return nil
}

// CreateUserReqValidationError is the validation error returned by
// CreateUserReq.Validate if the designated constraints aren't met.
type CreateUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserReqValidationError) ErrorName() string { return "CreateUserReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserReqValidationError{}

// Validate checks the field values on CreateUserReqInfo with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateUserReqInfo) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return CreateUserReqInfoValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if !_CreateUserReqInfo_Password_Pattern.MatchString(m.GetPassword()) {
		return CreateUserReqInfoValidationError{
			field:  "Password",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]{6,72}$\"",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 10 {
		return CreateUserReqInfoValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 10 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetIntroduction()); l < 1 || l > 1000 {
		return CreateUserReqInfoValidationError{
			field:  "Introduction",
			reason: "value length must be between 1 and 1000 runes, inclusive",
		}
	}

	if _, ok := Sex_name[int32(m.GetSex())]; !ok {
		return CreateUserReqInfoValidationError{
			field:  "Sex",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

func (m *CreateUserReqInfo) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *CreateUserReqInfo) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// CreateUserReqInfoValidationError is the validation error returned by
// CreateUserReqInfo.Validate if the designated constraints aren't met.
type CreateUserReqInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserReqInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserReqInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserReqInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserReqInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserReqInfoValidationError) ErrorName() string {
	return "CreateUserReqInfoValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserReqInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserReqInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserReqInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserReqInfoValidationError{}

var _CreateUserReqInfo_Password_Pattern = regexp.MustCompile("^[A-Za-z0-9]{6,72}$")

// Validate checks the field values on CreateUserRes with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateUserRes) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateUserResValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTokenPair()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateUserResValidationError{
				field:  "TokenPair",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateUserResValidationError is the validation error returned by
// CreateUserRes.Validate if the designated constraints aren't met.
type CreateUserResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserResValidationError) ErrorName() string { return "CreateUserResValidationError" }

// Error satisfies the builtin error interface
func (e CreateUserResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserResValidationError{}

// Validate checks the field values on UpdateUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdateUserReq) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Data.(type) {

	case *UpdateUserReq_Info:

		if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateUserReqValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *UpdateUserReq_ImageChunk:

		if len(m.GetImageChunk()) > 65536 {
			return UpdateUserReqValidationError{
				field:  "ImageChunk",
				reason: "value length must be at most 65536 bytes",
			}
		}

	default:
		return UpdateUserReqValidationError{
			field:  "Data",
			reason: "value is required",
		}

	}

	return nil
}

// UpdateUserReqValidationError is the validation error returned by
// UpdateUserReq.Validate if the designated constraints aren't met.
type UpdateUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserReqValidationError) ErrorName() string { return "UpdateUserReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserReqValidationError{}

// Validate checks the field values on UpdateUserReqInfo with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdateUserReqInfo) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return UpdateUserReqInfoValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 10 {
		return UpdateUserReqInfoValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 10 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetIntroduction()); l < 1 || l > 1000 {
		return UpdateUserReqInfoValidationError{
			field:  "Introduction",
			reason: "value length must be between 1 and 1000 runes, inclusive",
		}
	}

	return nil
}

func (m *UpdateUserReqInfo) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UpdateUserReqInfo) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UpdateUserReqInfoValidationError is the validation error returned by
// UpdateUserReqInfo.Validate if the designated constraints aren't met.
type UpdateUserReqInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserReqInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserReqInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserReqInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserReqInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserReqInfoValidationError) ErrorName() string {
	return "UpdateUserReqInfoValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserReqInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserReqInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserReqInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserReqInfoValidationError{}

// Validate checks the field values on UpdatePasswordReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdatePasswordReq) Validate() error {
	if m == nil {
		return nil
	}

	if !_UpdatePasswordReq_OldPassword_Pattern.MatchString(m.GetOldPassword()) {
		return UpdatePasswordReqValidationError{
			field:  "OldPassword",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]{6,72}$\"",
		}
	}

	if !_UpdatePasswordReq_NewPassword_Pattern.MatchString(m.GetNewPassword()) {
		return UpdatePasswordReqValidationError{
			field:  "NewPassword",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]{6,72}$\"",
		}
	}

	return nil
}

// UpdatePasswordReqValidationError is the validation error returned by
// UpdatePasswordReq.Validate if the designated constraints aren't met.
type UpdatePasswordReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePasswordReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePasswordReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePasswordReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePasswordReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePasswordReqValidationError) ErrorName() string {
	return "UpdatePasswordReqValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePasswordReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePasswordReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePasswordReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePasswordReqValidationError{}

var _UpdatePasswordReq_OldPassword_Pattern = regexp.MustCompile("^[A-Za-z0-9]{6,72}$")

var _UpdatePasswordReq_NewPassword_Pattern = regexp.MustCompile("^[A-Za-z0-9]{6,72}$")

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LoginReq) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return LoginReqValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if !_LoginReq_Password_Pattern.MatchString(m.GetPassword()) {
		return LoginReqValidationError{
			field:  "Password",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]{6,72}$\"",
		}
	}

	return nil
}

func (m *LoginReq) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *LoginReq) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

var _LoginReq_Password_Pattern = regexp.MustCompile("^[A-Za-z0-9]{6,72}$")

// Validate checks the field values on LoginRes with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LoginRes) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginResValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTokenPair()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginResValidationError{
				field:  "TokenPair",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LoginResValidationError is the validation error returned by
// LoginRes.Validate if the designated constraints aren't met.
type LoginResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResValidationError) ErrorName() string { return "LoginResValidationError" }

// Error satisfies the builtin error interface
func (e LoginResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResValidationError{}

// Validate checks the field values on LogoutReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LogoutReq) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// LogoutReqValidationError is the validation error returned by
// LogoutReq.Validate if the designated constraints aren't met.
type LogoutReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogoutReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogoutReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogoutReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogoutReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogoutReqValidationError) ErrorName() string { return "LogoutReqValidationError" }

// Error satisfies the builtin error interface
func (e LogoutReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogoutReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogoutReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogoutReqValidationError{}

// Validate checks the field values on RefreshIDTokenReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RefreshIDTokenReq) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RefreshIDTokenReqValidationError is the validation error returned by
// RefreshIDTokenReq.Validate if the designated constraints aren't met.
type RefreshIDTokenReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshIDTokenReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshIDTokenReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshIDTokenReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshIDTokenReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshIDTokenReqValidationError) ErrorName() string {
	return "RefreshIDTokenReqValidationError"
}

// Error satisfies the builtin error interface
func (e RefreshIDTokenReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshIDTokenReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshIDTokenReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshIDTokenReqValidationError{}

// Validate checks the field values on RefreshIDTokenRes with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RefreshIDTokenRes) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTokenPair()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RefreshIDTokenResValidationError{
				field:  "TokenPair",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RefreshIDTokenResValidationError is the validation error returned by
// RefreshIDTokenRes.Validate if the designated constraints aren't met.
type RefreshIDTokenResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshIDTokenResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshIDTokenResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshIDTokenResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshIDTokenResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshIDTokenResValidationError) ErrorName() string {
	return "RefreshIDTokenResValidationError"
}

// Error satisfies the builtin error interface
func (e RefreshIDTokenResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshIDTokenRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshIDTokenResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshIDTokenResValidationError{}
