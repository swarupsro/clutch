// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chaos/experimentation/v1/experimentation.proto

package experimentationv1

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
var _experimentation_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Experiment with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Experiment) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetTestConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExperimentValidationError{
				field:  "TestConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExperimentValidationError is the validation error returned by
// Experiment.Validate if the designated constraints aren't met.
type ExperimentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExperimentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExperimentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExperimentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExperimentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExperimentValidationError) ErrorName() string { return "ExperimentValidationError" }

// Error satisfies the builtin error interface
func (e ExperimentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExperiment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExperimentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExperimentValidationError{}

// Validate checks the field values on CreateExperimentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateExperimentsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetExperiments()) < 1 {
		return CreateExperimentsRequestValidationError{
			field:  "Experiments",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetExperiments() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateExperimentsRequestValidationError{
					field:  fmt.Sprintf("Experiments[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateExperimentsRequestValidationError is the validation error returned by
// CreateExperimentsRequest.Validate if the designated constraints aren't met.
type CreateExperimentsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateExperimentsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateExperimentsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateExperimentsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateExperimentsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateExperimentsRequestValidationError) ErrorName() string {
	return "CreateExperimentsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateExperimentsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateExperimentsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateExperimentsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateExperimentsRequestValidationError{}

// Validate checks the field values on CreateExperimentsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateExperimentsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetExperiments() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateExperimentsResponseValidationError{
					field:  fmt.Sprintf("Experiments[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateExperimentsResponseValidationError is the validation error returned by
// CreateExperimentsResponse.Validate if the designated constraints aren't met.
type CreateExperimentsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateExperimentsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateExperimentsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateExperimentsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateExperimentsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateExperimentsResponseValidationError) ErrorName() string {
	return "CreateExperimentsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateExperimentsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateExperimentsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateExperimentsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateExperimentsResponseValidationError{}

// Validate checks the field values on GetExperimentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetExperimentsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Convert

	return nil
}

// GetExperimentsRequestValidationError is the validation error returned by
// GetExperimentsRequest.Validate if the designated constraints aren't met.
type GetExperimentsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExperimentsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExperimentsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExperimentsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExperimentsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExperimentsRequestValidationError) ErrorName() string {
	return "GetExperimentsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetExperimentsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExperimentsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExperimentsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExperimentsRequestValidationError{}

// Validate checks the field values on GetExperimentsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetExperimentsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetExperiments() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetExperimentsResponseValidationError{
					field:  fmt.Sprintf("Experiments[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetExperimentsResponseValidationError is the validation error returned by
// GetExperimentsResponse.Validate if the designated constraints aren't met.
type GetExperimentsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExperimentsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExperimentsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExperimentsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExperimentsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExperimentsResponseValidationError) ErrorName() string {
	return "GetExperimentsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetExperimentsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExperimentsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExperimentsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExperimentsResponseValidationError{}

// Validate checks the field values on DeleteExperimentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteExperimentsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteExperimentsRequestValidationError is the validation error returned by
// DeleteExperimentsRequest.Validate if the designated constraints aren't met.
type DeleteExperimentsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteExperimentsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteExperimentsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteExperimentsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteExperimentsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteExperimentsRequestValidationError) ErrorName() string {
	return "DeleteExperimentsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteExperimentsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteExperimentsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteExperimentsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteExperimentsRequestValidationError{}

// Validate checks the field values on DeleteExperimentsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteExperimentsResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteExperimentsResponseValidationError is the validation error returned by
// DeleteExperimentsResponse.Validate if the designated constraints aren't met.
type DeleteExperimentsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteExperimentsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteExperimentsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteExperimentsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteExperimentsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteExperimentsResponseValidationError) ErrorName() string {
	return "DeleteExperimentsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteExperimentsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteExperimentsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteExperimentsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteExperimentsResponseValidationError{}
