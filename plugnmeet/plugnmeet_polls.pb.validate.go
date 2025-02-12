// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: plugnmeet_polls.proto

package plugnmeet

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

// Validate checks the field values on CreatePollReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreatePollReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePollReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreatePollReqMultiError, or
// nil if none found.
func (m *CreatePollReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePollReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	// no validation rules for UserId

	// no validation rules for PollId

	// no validation rules for Question

	for idx, item := range m.GetOptions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreatePollReqValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreatePollReqValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreatePollReqValidationError{
					field:  fmt.Sprintf("Options[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreatePollReqMultiError(errors)
	}

	return nil
}

// CreatePollReqMultiError is an error wrapping multiple validation errors
// returned by CreatePollReq.ValidateAll() if the designated constraints
// aren't met.
type CreatePollReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePollReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePollReqMultiError) AllErrors() []error { return m }

// CreatePollReqValidationError is the validation error returned by
// CreatePollReq.Validate if the designated constraints aren't met.
type CreatePollReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePollReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePollReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePollReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePollReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePollReqValidationError) ErrorName() string { return "CreatePollReqValidationError" }

// Error satisfies the builtin error interface
func (e CreatePollReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePollReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePollReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePollReqValidationError{}

// Validate checks the field values on CreatePollOptions with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreatePollOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePollOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePollOptionsMultiError, or nil if none found.
func (m *CreatePollOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePollOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	if len(errors) > 0 {
		return CreatePollOptionsMultiError(errors)
	}

	return nil
}

// CreatePollOptionsMultiError is an error wrapping multiple validation errors
// returned by CreatePollOptions.ValidateAll() if the designated constraints
// aren't met.
type CreatePollOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePollOptionsMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePollOptionsMultiError) AllErrors() []error { return m }

// CreatePollOptionsValidationError is the validation error returned by
// CreatePollOptions.Validate if the designated constraints aren't met.
type CreatePollOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePollOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePollOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePollOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePollOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePollOptionsValidationError) ErrorName() string {
	return "CreatePollOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePollOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePollOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePollOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePollOptionsValidationError{}

// Validate checks the field values on PollInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PollInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PollInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PollInfoMultiError, or nil
// if none found.
func (m *PollInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *PollInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for RoomId

	// no validation rules for Question

	for idx, item := range m.GetOptions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PollInfoValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PollInfoValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PollInfoValidationError{
					field:  fmt.Sprintf("Options[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for IsRunning

	// no validation rules for Created

	// no validation rules for CreatedBy

	// no validation rules for ClosedBy

	if len(errors) > 0 {
		return PollInfoMultiError(errors)
	}

	return nil
}

// PollInfoMultiError is an error wrapping multiple validation errors returned
// by PollInfo.ValidateAll() if the designated constraints aren't met.
type PollInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PollInfoMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PollInfoMultiError) AllErrors() []error { return m }

// PollInfoValidationError is the validation error returned by
// PollInfo.Validate if the designated constraints aren't met.
type PollInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PollInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PollInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PollInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PollInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PollInfoValidationError) ErrorName() string { return "PollInfoValidationError" }

// Error satisfies the builtin error interface
func (e PollInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPollInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PollInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PollInfoValidationError{}

// Validate checks the field values on SubmitPollResponseReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SubmitPollResponseReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubmitPollResponseReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubmitPollResponseReqMultiError, or nil if none found.
func (m *SubmitPollResponseReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SubmitPollResponseReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	// no validation rules for UserId

	// no validation rules for Name

	// no validation rules for PollId

	// no validation rules for SelectedOption

	if len(errors) > 0 {
		return SubmitPollResponseReqMultiError(errors)
	}

	return nil
}

// SubmitPollResponseReqMultiError is an error wrapping multiple validation
// errors returned by SubmitPollResponseReq.ValidateAll() if the designated
// constraints aren't met.
type SubmitPollResponseReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubmitPollResponseReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubmitPollResponseReqMultiError) AllErrors() []error { return m }

// SubmitPollResponseReqValidationError is the validation error returned by
// SubmitPollResponseReq.Validate if the designated constraints aren't met.
type SubmitPollResponseReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitPollResponseReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitPollResponseReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitPollResponseReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitPollResponseReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitPollResponseReqValidationError) ErrorName() string {
	return "SubmitPollResponseReqValidationError"
}

// Error satisfies the builtin error interface
func (e SubmitPollResponseReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitPollResponseReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitPollResponseReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitPollResponseReqValidationError{}

// Validate checks the field values on ClosePollReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClosePollReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClosePollReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClosePollReqMultiError, or
// nil if none found.
func (m *ClosePollReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ClosePollReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	// no validation rules for UserId

	// no validation rules for PollId

	if len(errors) > 0 {
		return ClosePollReqMultiError(errors)
	}

	return nil
}

// ClosePollReqMultiError is an error wrapping multiple validation errors
// returned by ClosePollReq.ValidateAll() if the designated constraints aren't met.
type ClosePollReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClosePollReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClosePollReqMultiError) AllErrors() []error { return m }

// ClosePollReqValidationError is the validation error returned by
// ClosePollReq.Validate if the designated constraints aren't met.
type ClosePollReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClosePollReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClosePollReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClosePollReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClosePollReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClosePollReqValidationError) ErrorName() string { return "ClosePollReqValidationError" }

// Error satisfies the builtin error interface
func (e ClosePollReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClosePollReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClosePollReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClosePollReqValidationError{}

// Validate checks the field values on PollResponsesResultOptions with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PollResponsesResultOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PollResponsesResultOptions with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PollResponsesResultOptionsMultiError, or nil if none found.
func (m *PollResponsesResultOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *PollResponsesResultOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	// no validation rules for VoteCount

	if len(errors) > 0 {
		return PollResponsesResultOptionsMultiError(errors)
	}

	return nil
}

// PollResponsesResultOptionsMultiError is an error wrapping multiple
// validation errors returned by PollResponsesResultOptions.ValidateAll() if
// the designated constraints aren't met.
type PollResponsesResultOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PollResponsesResultOptionsMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PollResponsesResultOptionsMultiError) AllErrors() []error { return m }

// PollResponsesResultOptionsValidationError is the validation error returned
// by PollResponsesResultOptions.Validate if the designated constraints aren't met.
type PollResponsesResultOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PollResponsesResultOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PollResponsesResultOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PollResponsesResultOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PollResponsesResultOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PollResponsesResultOptionsValidationError) ErrorName() string {
	return "PollResponsesResultOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e PollResponsesResultOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPollResponsesResultOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PollResponsesResultOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PollResponsesResultOptionsValidationError{}

// Validate checks the field values on PollResponsesResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PollResponsesResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PollResponsesResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PollResponsesResultMultiError, or nil if none found.
func (m *PollResponsesResult) ValidateAll() error {
	return m.validate(true)
}

func (m *PollResponsesResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Question

	// no validation rules for TotalResponses

	for idx, item := range m.GetOptions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PollResponsesResultValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PollResponsesResultValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PollResponsesResultValidationError{
					field:  fmt.Sprintf("Options[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PollResponsesResultMultiError(errors)
	}

	return nil
}

// PollResponsesResultMultiError is an error wrapping multiple validation
// errors returned by PollResponsesResult.ValidateAll() if the designated
// constraints aren't met.
type PollResponsesResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PollResponsesResultMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PollResponsesResultMultiError) AllErrors() []error { return m }

// PollResponsesResultValidationError is the validation error returned by
// PollResponsesResult.Validate if the designated constraints aren't met.
type PollResponsesResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PollResponsesResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PollResponsesResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PollResponsesResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PollResponsesResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PollResponsesResultValidationError) ErrorName() string {
	return "PollResponsesResultValidationError"
}

// Error satisfies the builtin error interface
func (e PollResponsesResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPollResponsesResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PollResponsesResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PollResponsesResultValidationError{}

// Validate checks the field values on PollsStats with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PollsStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PollsStats with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PollsStatsMultiError, or
// nil if none found.
func (m *PollsStats) ValidateAll() error {
	return m.validate(true)
}

func (m *PollsStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalPolls

	// no validation rules for TotalRunning

	if len(errors) > 0 {
		return PollsStatsMultiError(errors)
	}

	return nil
}

// PollsStatsMultiError is an error wrapping multiple validation errors
// returned by PollsStats.ValidateAll() if the designated constraints aren't met.
type PollsStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PollsStatsMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PollsStatsMultiError) AllErrors() []error { return m }

// PollsStatsValidationError is the validation error returned by
// PollsStats.Validate if the designated constraints aren't met.
type PollsStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PollsStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PollsStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PollsStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PollsStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PollsStatsValidationError) ErrorName() string { return "PollsStatsValidationError" }

// Error satisfies the builtin error interface
func (e PollsStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPollsStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PollsStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PollsStatsValidationError{}

// Validate checks the field values on PollResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PollResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PollResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PollResponseMultiError, or
// nil if none found.
func (m *PollResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PollResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	// no validation rules for Responses

	for idx, item := range m.GetPolls() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PollResponseValidationError{
						field:  fmt.Sprintf("Polls[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PollResponseValidationError{
						field:  fmt.Sprintf("Polls[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PollResponseValidationError{
					field:  fmt.Sprintf("Polls[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.PollId != nil {
		// no validation rules for PollId
	}

	if m.TotalResponses != nil {
		// no validation rules for TotalResponses
	}

	if m.Voted != nil {
		// no validation rules for Voted
	}

	if m.Stats != nil {

		if all {
			switch v := interface{}(m.GetStats()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PollResponseValidationError{
						field:  "Stats",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PollResponseValidationError{
						field:  "Stats",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStats()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PollResponseValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.TotalPolls != nil {
		// no validation rules for TotalPolls
	}

	if m.TotalRunning != nil {
		// no validation rules for TotalRunning
	}

	if m.PollResponsesResult != nil {

		if all {
			switch v := interface{}(m.GetPollResponsesResult()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PollResponseValidationError{
						field:  "PollResponsesResult",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PollResponseValidationError{
						field:  "PollResponsesResult",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPollResponsesResult()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PollResponseValidationError{
					field:  "PollResponsesResult",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PollResponseMultiError(errors)
	}

	return nil
}

// PollResponseMultiError is an error wrapping multiple validation errors
// returned by PollResponse.ValidateAll() if the designated constraints aren't met.
type PollResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PollResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PollResponseMultiError) AllErrors() []error { return m }

// PollResponseValidationError is the validation error returned by
// PollResponse.Validate if the designated constraints aren't met.
type PollResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PollResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PollResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PollResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PollResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PollResponseValidationError) ErrorName() string { return "PollResponseValidationError" }

// Error satisfies the builtin error interface
func (e PollResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPollResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PollResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PollResponseValidationError{}
