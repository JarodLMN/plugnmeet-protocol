// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: plugnmeet_common.proto

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

// Validate checks the field values on CommonNotifyEvent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CommonNotifyEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonNotifyEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonNotifyEventMultiError, or nil if none found.
func (m *CommonNotifyEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonNotifyEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Event != nil {
		// no validation rules for Event
	}

	if m.Room != nil {

		if all {
			switch v := interface{}(m.GetRoom()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "Room",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "Room",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRoom()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonNotifyEventValidationError{
					field:  "Room",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Participant != nil {

		if all {
			switch v := interface{}(m.GetParticipant()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "Participant",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "Participant",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetParticipant()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonNotifyEventValidationError{
					field:  "Participant",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.RecordingInfo != nil {

		if all {
			switch v := interface{}(m.GetRecordingInfo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "RecordingInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "RecordingInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRecordingInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonNotifyEventValidationError{
					field:  "RecordingInfo",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.SpeechService != nil {

		if all {
			switch v := interface{}(m.GetSpeechService()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "SpeechService",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "SpeechService",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSpeechService()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonNotifyEventValidationError{
					field:  "SpeechService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Track != nil {

		if all {
			switch v := interface{}(m.GetTrack()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "Track",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "Track",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTrack()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonNotifyEventValidationError{
					field:  "Track",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Analytics != nil {

		if all {
			switch v := interface{}(m.GetAnalytics()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "Analytics",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonNotifyEventValidationError{
						field:  "Analytics",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAnalytics()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonNotifyEventValidationError{
					field:  "Analytics",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.CreatedAt != nil {
		// no validation rules for CreatedAt
	}

	if len(errors) > 0 {
		return CommonNotifyEventMultiError(errors)
	}

	return nil
}

// CommonNotifyEventMultiError is an error wrapping multiple validation errors
// returned by CommonNotifyEvent.ValidateAll() if the designated constraints
// aren't met.
type CommonNotifyEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonNotifyEventMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonNotifyEventMultiError) AllErrors() []error { return m }

// CommonNotifyEventValidationError is the validation error returned by
// CommonNotifyEvent.Validate if the designated constraints aren't met.
type CommonNotifyEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonNotifyEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonNotifyEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonNotifyEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonNotifyEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonNotifyEventValidationError) ErrorName() string {
	return "CommonNotifyEventValidationError"
}

// Error satisfies the builtin error interface
func (e CommonNotifyEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonNotifyEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonNotifyEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonNotifyEventValidationError{}

// Validate checks the field values on NotifyEventRoom with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *NotifyEventRoom) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotifyEventRoom with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotifyEventRoomMultiError, or nil if none found.
func (m *NotifyEventRoom) ValidateAll() error {
	return m.validate(true)
}

func (m *NotifyEventRoom) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEnabledCodecs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NotifyEventRoomValidationError{
						field:  fmt.Sprintf("EnabledCodecs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NotifyEventRoomValidationError{
						field:  fmt.Sprintf("EnabledCodecs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NotifyEventRoomValidationError{
					field:  fmt.Sprintf("EnabledCodecs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Sid != nil {
		// no validation rules for Sid
	}

	if m.RoomId != nil {
		// no validation rules for RoomId
	}

	if m.EmptyTimeout != nil {
		// no validation rules for EmptyTimeout
	}

	if m.MaxParticipants != nil {
		// no validation rules for MaxParticipants
	}

	if m.CreationTime != nil {
		// no validation rules for CreationTime
	}

	if m.Metadata != nil {
		// no validation rules for Metadata
	}

	if m.NumParticipants != nil {
		// no validation rules for NumParticipants
	}

	if len(errors) > 0 {
		return NotifyEventRoomMultiError(errors)
	}

	return nil
}

// NotifyEventRoomMultiError is an error wrapping multiple validation errors
// returned by NotifyEventRoom.ValidateAll() if the designated constraints
// aren't met.
type NotifyEventRoomMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotifyEventRoomMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotifyEventRoomMultiError) AllErrors() []error { return m }

// NotifyEventRoomValidationError is the validation error returned by
// NotifyEventRoom.Validate if the designated constraints aren't met.
type NotifyEventRoomValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotifyEventRoomValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotifyEventRoomValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotifyEventRoomValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotifyEventRoomValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotifyEventRoomValidationError) ErrorName() string { return "NotifyEventRoomValidationError" }

// Error satisfies the builtin error interface
func (e NotifyEventRoomValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotifyEventRoom.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotifyEventRoomValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotifyEventRoomValidationError{}

// Validate checks the field values on RecordingInfoEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RecordingInfoEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RecordingInfoEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RecordingInfoEventMultiError, or nil if none found.
func (m *RecordingInfoEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *RecordingInfoEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RecordId

	// no validation rules for RecorderId

	// no validation rules for RecorderMsg

	if m.FilePath != nil {
		// no validation rules for FilePath
	}

	if m.FileSize != nil {
		// no validation rules for FileSize
	}

	if len(errors) > 0 {
		return RecordingInfoEventMultiError(errors)
	}

	return nil
}

// RecordingInfoEventMultiError is an error wrapping multiple validation errors
// returned by RecordingInfoEvent.ValidateAll() if the designated constraints
// aren't met.
type RecordingInfoEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RecordingInfoEventMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RecordingInfoEventMultiError) AllErrors() []error { return m }

// RecordingInfoEventValidationError is the validation error returned by
// RecordingInfoEvent.Validate if the designated constraints aren't met.
type RecordingInfoEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RecordingInfoEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RecordingInfoEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RecordingInfoEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RecordingInfoEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RecordingInfoEventValidationError) ErrorName() string {
	return "RecordingInfoEventValidationError"
}

// Error satisfies the builtin error interface
func (e RecordingInfoEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRecordingInfoEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RecordingInfoEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RecordingInfoEventValidationError{}

// Validate checks the field values on SpeechServiceEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SpeechServiceEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SpeechServiceEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SpeechServiceEventMultiError, or nil if none found.
func (m *SpeechServiceEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *SpeechServiceEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalUsage

	if m.UserId != nil {
		// no validation rules for UserId
	}

	if len(errors) > 0 {
		return SpeechServiceEventMultiError(errors)
	}

	return nil
}

// SpeechServiceEventMultiError is an error wrapping multiple validation errors
// returned by SpeechServiceEvent.ValidateAll() if the designated constraints
// aren't met.
type SpeechServiceEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SpeechServiceEventMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SpeechServiceEventMultiError) AllErrors() []error { return m }

// SpeechServiceEventValidationError is the validation error returned by
// SpeechServiceEvent.Validate if the designated constraints aren't met.
type SpeechServiceEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpeechServiceEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpeechServiceEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpeechServiceEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpeechServiceEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpeechServiceEventValidationError) ErrorName() string {
	return "SpeechServiceEventValidationError"
}

// Error satisfies the builtin error interface
func (e SpeechServiceEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpeechServiceEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpeechServiceEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpeechServiceEventValidationError{}

// Validate checks the field values on AnalyticsEvent with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AnalyticsEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AnalyticsEvent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AnalyticsEventMultiError,
// or nil if none found.
func (m *AnalyticsEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *AnalyticsEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.FileId != nil {
		// no validation rules for FileId
	}

	if len(errors) > 0 {
		return AnalyticsEventMultiError(errors)
	}

	return nil
}

// AnalyticsEventMultiError is an error wrapping multiple validation errors
// returned by AnalyticsEvent.ValidateAll() if the designated constraints
// aren't met.
type AnalyticsEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AnalyticsEventMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AnalyticsEventMultiError) AllErrors() []error { return m }

// AnalyticsEventValidationError is the validation error returned by
// AnalyticsEvent.Validate if the designated constraints aren't met.
type AnalyticsEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AnalyticsEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AnalyticsEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AnalyticsEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AnalyticsEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AnalyticsEventValidationError) ErrorName() string { return "AnalyticsEventValidationError" }

// Error satisfies the builtin error interface
func (e AnalyticsEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAnalyticsEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AnalyticsEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AnalyticsEventValidationError{}
