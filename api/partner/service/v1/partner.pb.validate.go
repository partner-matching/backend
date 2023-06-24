// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: partner/service/v1/partner.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetTeamResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTeamResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTeamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTeamResponseMultiError, or nil if none found.
func (m *GetTeamResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTeamResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTeamResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTeamResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTeamResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetTeamResponseMultiError(errors)
	}

	return nil
}

// GetTeamResponseMultiError is an error wrapping multiple validation errors
// returned by GetTeamResponse.ValidateAll() if the designated constraints
// aren't met.
type GetTeamResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTeamResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTeamResponseMultiError) AllErrors() []error { return m }

// GetTeamResponseValidationError is the validation error returned by
// GetTeamResponse.Validate if the designated constraints aren't met.
type GetTeamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTeamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTeamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTeamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTeamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTeamResponseValidationError) ErrorName() string { return "GetTeamResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetTeamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTeamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTeamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTeamResponseValidationError{}

// Validate checks the field values on GetTeamListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTeamListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTeamListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTeamListResponseMultiError, or nil if none found.
func (m *GetTeamListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTeamListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTeamListResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTeamListResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTeamListResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTeamListResponseMultiError(errors)
	}

	return nil
}

// GetTeamListResponseMultiError is an error wrapping multiple validation
// errors returned by GetTeamListResponse.ValidateAll() if the designated
// constraints aren't met.
type GetTeamListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTeamListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTeamListResponseMultiError) AllErrors() []error { return m }

// GetTeamListResponseValidationError is the validation error returned by
// GetTeamListResponse.Validate if the designated constraints aren't met.
type GetTeamListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTeamListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTeamListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTeamListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTeamListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTeamListResponseValidationError) ErrorName() string {
	return "GetTeamListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTeamListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTeamListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTeamListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTeamListResponseValidationError{}

// Validate checks the field values on DeleteTeamReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteTeamReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTeamReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteTeamReqMultiError, or
// nil if none found.
func (m *DeleteTeamReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTeamReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteTeamReqMultiError(errors)
	}

	return nil
}

// DeleteTeamReqMultiError is an error wrapping multiple validation errors
// returned by DeleteTeamReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteTeamReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTeamReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTeamReqMultiError) AllErrors() []error { return m }

// DeleteTeamReqValidationError is the validation error returned by
// DeleteTeamReq.Validate if the designated constraints aren't met.
type DeleteTeamReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTeamReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTeamReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTeamReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTeamReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTeamReqValidationError) ErrorName() string { return "DeleteTeamReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteTeamReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTeamReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTeamReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTeamReqValidationError{}

// Validate checks the field values on UpdateTeamReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateTeamReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTeamReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateTeamReqMultiError, or
// nil if none found.
func (m *UpdateTeamReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTeamReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for ExpireTime

	// no validation rules for Status

	// no validation rules for Password

	// no validation rules for Description

	if len(errors) > 0 {
		return UpdateTeamReqMultiError(errors)
	}

	return nil
}

// UpdateTeamReqMultiError is an error wrapping multiple validation errors
// returned by UpdateTeamReq.ValidateAll() if the designated constraints
// aren't met.
type UpdateTeamReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTeamReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTeamReqMultiError) AllErrors() []error { return m }

// UpdateTeamReqValidationError is the validation error returned by
// UpdateTeamReq.Validate if the designated constraints aren't met.
type UpdateTeamReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTeamReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTeamReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTeamReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTeamReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTeamReqValidationError) ErrorName() string { return "UpdateTeamReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateTeamReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTeamReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTeamReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTeamReqValidationError{}

// Validate checks the field values on GetTeamReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTeamReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTeamReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTeamReqMultiError, or
// nil if none found.
func (m *GetTeamReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTeamReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetTeamReqMultiError(errors)
	}

	return nil
}

// GetTeamReqMultiError is an error wrapping multiple validation errors
// returned by GetTeamReq.ValidateAll() if the designated constraints aren't met.
type GetTeamReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTeamReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTeamReqMultiError) AllErrors() []error { return m }

// GetTeamReqValidationError is the validation error returned by
// GetTeamReq.Validate if the designated constraints aren't met.
type GetTeamReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTeamReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTeamReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTeamReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTeamReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTeamReqValidationError) ErrorName() string { return "GetTeamReqValidationError" }

// Error satisfies the builtin error interface
func (e GetTeamReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTeamReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTeamReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTeamReqValidationError{}

// Validate checks the field values on GetTeamListReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTeamListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTeamListReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTeamListReqMultiError,
// or nil if none found.
func (m *GetTeamListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTeamListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if all {
		switch v := interface{}(m.GetQuery()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTeamListReqValidationError{
					field:  "Query",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTeamListReqValidationError{
					field:  "Query",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuery()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTeamListReqValidationError{
				field:  "Query",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetTeamListReqMultiError(errors)
	}

	return nil
}

// GetTeamListReqMultiError is an error wrapping multiple validation errors
// returned by GetTeamListReq.ValidateAll() if the designated constraints
// aren't met.
type GetTeamListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTeamListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTeamListReqMultiError) AllErrors() []error { return m }

// GetTeamListReqValidationError is the validation error returned by
// GetTeamListReq.Validate if the designated constraints aren't met.
type GetTeamListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTeamListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTeamListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTeamListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTeamListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTeamListReqValidationError) ErrorName() string { return "GetTeamListReqValidationError" }

// Error satisfies the builtin error interface
func (e GetTeamListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTeamListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTeamListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTeamListReqValidationError{}

// Validate checks the field values on JoinTeamReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *JoinTeamReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JoinTeamReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in JoinTeamReqMultiError, or
// nil if none found.
func (m *JoinTeamReq) ValidateAll() error {
	return m.validate(true)
}

func (m *JoinTeamReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for PassWord

	if len(errors) > 0 {
		return JoinTeamReqMultiError(errors)
	}

	return nil
}

// JoinTeamReqMultiError is an error wrapping multiple validation errors
// returned by JoinTeamReq.ValidateAll() if the designated constraints aren't met.
type JoinTeamReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JoinTeamReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JoinTeamReqMultiError) AllErrors() []error { return m }

// JoinTeamReqValidationError is the validation error returned by
// JoinTeamReq.Validate if the designated constraints aren't met.
type JoinTeamReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JoinTeamReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JoinTeamReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JoinTeamReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JoinTeamReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JoinTeamReqValidationError) ErrorName() string { return "JoinTeamReqValidationError" }

// Error satisfies the builtin error interface
func (e JoinTeamReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJoinTeamReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JoinTeamReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JoinTeamReqValidationError{}

// Validate checks the field values on QuitTeamReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuitTeamReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuitTeamReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuitTeamReqMultiError, or
// nil if none found.
func (m *QuitTeamReq) ValidateAll() error {
	return m.validate(true)
}

func (m *QuitTeamReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return QuitTeamReqMultiError(errors)
	}

	return nil
}

// QuitTeamReqMultiError is an error wrapping multiple validation errors
// returned by QuitTeamReq.ValidateAll() if the designated constraints aren't met.
type QuitTeamReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuitTeamReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuitTeamReqMultiError) AllErrors() []error { return m }

// QuitTeamReqValidationError is the validation error returned by
// QuitTeamReq.Validate if the designated constraints aren't met.
type QuitTeamReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuitTeamReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuitTeamReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuitTeamReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuitTeamReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuitTeamReqValidationError) ErrorName() string { return "QuitTeamReqValidationError" }

// Error satisfies the builtin error interface
func (e QuitTeamReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuitTeamReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuitTeamReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuitTeamReqValidationError{}

// Validate checks the field values on Team with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Team) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Team with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TeamMultiError, or nil if none found.
func (m *Team) ValidateAll() error {
	return m.validate(true)
}

func (m *Team) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for MaxNum

	// no validation rules for ExpireTime

	// no validation rules for UserId

	// no validation rules for Status

	// no validation rules for Password

	// no validation rules for Description

	if len(errors) > 0 {
		return TeamMultiError(errors)
	}

	return nil
}

// TeamMultiError is an error wrapping multiple validation errors returned by
// Team.ValidateAll() if the designated constraints aren't met.
type TeamMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TeamMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TeamMultiError) AllErrors() []error { return m }

// TeamValidationError is the validation error returned by Team.Validate if the
// designated constraints aren't met.
type TeamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TeamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TeamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TeamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TeamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TeamValidationError) ErrorName() string { return "TeamValidationError" }

// Error satisfies the builtin error interface
func (e TeamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTeam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TeamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TeamValidationError{}

// Validate checks the field values on GetTeamListResponse_TeamInfo with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTeamListResponse_TeamInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTeamListResponse_TeamInfo with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTeamListResponse_TeamInfoMultiError, or nil if none found.
func (m *GetTeamListResponse_TeamInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTeamListResponse_TeamInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTeam()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTeamListResponse_TeamInfoValidationError{
					field:  "Team",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTeamListResponse_TeamInfoValidationError{
					field:  "Team",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTeam()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTeamListResponse_TeamInfoValidationError{
				field:  "Team",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUserInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTeamListResponse_TeamInfoValidationError{
					field:  "UserInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTeamListResponse_TeamInfoValidationError{
					field:  "UserInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTeamListResponse_TeamInfoValidationError{
				field:  "UserInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetTeamListResponse_TeamInfoMultiError(errors)
	}

	return nil
}

// GetTeamListResponse_TeamInfoMultiError is an error wrapping multiple
// validation errors returned by GetTeamListResponse_TeamInfo.ValidateAll() if
// the designated constraints aren't met.
type GetTeamListResponse_TeamInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTeamListResponse_TeamInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTeamListResponse_TeamInfoMultiError) AllErrors() []error { return m }

// GetTeamListResponse_TeamInfoValidationError is the validation error returned
// by GetTeamListResponse_TeamInfo.Validate if the designated constraints
// aren't met.
type GetTeamListResponse_TeamInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTeamListResponse_TeamInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTeamListResponse_TeamInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTeamListResponse_TeamInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTeamListResponse_TeamInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTeamListResponse_TeamInfoValidationError) ErrorName() string {
	return "GetTeamListResponse_TeamInfoValidationError"
}

// Error satisfies the builtin error interface
func (e GetTeamListResponse_TeamInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTeamListResponse_TeamInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTeamListResponse_TeamInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTeamListResponse_TeamInfoValidationError{}
