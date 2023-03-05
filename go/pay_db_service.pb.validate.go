// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pay_db_service.proto

package pay

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
var _pay_db_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on BasePayOrder with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BasePayOrder) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Version

	// no validation rules for OutTradeNo

	// no validation rules for ChannelAccount

	// no validation rules for ChannelOrderId

	// no validation rules for GatewayOrderId

	// no validation rules for PayAmount

	// no validation rules for Currency

	// no validation rules for NotifyUrl

	// no validation rules for ReturnUrl

	// no validation rules for AppId

	// no validation rules for SignType

	// no validation rules for OrderTime

	// no validation rules for RequestTime

	// no validation rules for CreateDate

	// no validation rules for UserIp

	// no validation rules for UserId

	// no validation rules for PayerAccount

	// no validation rules for ProductId

	// no validation rules for ProductName

	// no validation rules for ProductDescribe

	// no validation rules for CallbackJson

	// no validation rules for ExtJson

	// no validation rules for ChannelResponseJson

	// no validation rules for ErrorMessage

	// no validation rules for ChannelId

	// no validation rules for Method

	// no validation rules for Remark

	return nil
}

// BasePayOrderValidationError is the validation error returned by
// BasePayOrder.Validate if the designated constraints aren't met.
type BasePayOrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BasePayOrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BasePayOrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BasePayOrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BasePayOrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BasePayOrderValidationError) ErrorName() string { return "BasePayOrderValidationError" }

// Error satisfies the builtin error interface
func (e BasePayOrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBasePayOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BasePayOrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BasePayOrderValidationError{}

// Validate checks the field values on PayOrder with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PayOrder) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetBasePayOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PayOrderValidationError{
				field:  "BasePayOrder",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OrderStatus

	return nil
}

// PayOrderValidationError is the validation error returned by
// PayOrder.Validate if the designated constraints aren't met.
type PayOrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayOrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayOrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayOrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayOrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayOrderValidationError) ErrorName() string { return "PayOrderValidationError" }

// Error satisfies the builtin error interface
func (e PayOrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayOrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayOrderValidationError{}

// Validate checks the field values on PayOrderOk with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PayOrderOk) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetBasePayOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PayOrderOkValidationError{
				field:  "BasePayOrder",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SuccessTime

	// no validation rules for BalanceDate

	// no validation rules for FareAmt

	// no validation rules for FactAmt

	// no validation rules for SendNotifyStats

	return nil
}

// PayOrderOkValidationError is the validation error returned by
// PayOrderOk.Validate if the designated constraints aren't met.
type PayOrderOkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayOrderOkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayOrderOkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayOrderOkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayOrderOkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayOrderOkValidationError) ErrorName() string { return "PayOrderOkValidationError" }

// Error satisfies the builtin error interface
func (e PayOrderOkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayOrderOk.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayOrderOkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayOrderOkValidationError{}

// Validate checks the field values on PayNotify with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PayNotify) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for GatewayOrderId

	// no validation rules for CreateDate

	// no validation rules for FailTimes

	// no validation rules for NotifyTime

	// no validation rules for Status

	// no validation rules for NextNotifyTime

	// no validation rules for ErrorMessage

	return nil
}

// PayNotifyValidationError is the validation error returned by
// PayNotify.Validate if the designated constraints aren't met.
type PayNotifyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayNotifyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayNotifyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayNotifyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayNotifyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayNotifyValidationError) ErrorName() string { return "PayNotifyValidationError" }

// Error satisfies the builtin error interface
func (e PayNotifyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayNotify.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayNotifyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayNotifyValidationError{}

// Validate checks the field values on PayNotifyOk with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PayNotifyOk) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for GatewayOrderId

	// no validation rules for CreateDate

	// no validation rules for FailTimes

	// no validation rules for NotifyTime

	return nil
}

// PayNotifyOkValidationError is the validation error returned by
// PayNotifyOk.Validate if the designated constraints aren't met.
type PayNotifyOkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayNotifyOkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayNotifyOkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayNotifyOkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayNotifyOkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayNotifyOkValidationError) ErrorName() string { return "PayNotifyOkValidationError" }

// Error satisfies the builtin error interface
func (e PayNotifyOkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayNotifyOk.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayNotifyOkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayNotifyOkValidationError{}

// Validate checks the field values on PayOrderResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PayOrderResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPayOrders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PayOrderResponseValidationError{
					field:  fmt.Sprintf("PayOrders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PayOrderResponseValidationError is the validation error returned by
// PayOrderResponse.Validate if the designated constraints aren't met.
type PayOrderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayOrderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayOrderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayOrderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayOrderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayOrderResponseValidationError) ErrorName() string { return "PayOrderResponseValidationError" }

// Error satisfies the builtin error interface
func (e PayOrderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayOrderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayOrderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayOrderResponseValidationError{}

// Validate checks the field values on PayOrderOkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PayOrderOkResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPayOrderOks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PayOrderOkResponseValidationError{
					field:  fmt.Sprintf("PayOrderOks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PayOrderOkResponseValidationError is the validation error returned by
// PayOrderOkResponse.Validate if the designated constraints aren't met.
type PayOrderOkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayOrderOkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayOrderOkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayOrderOkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayOrderOkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayOrderOkResponseValidationError) ErrorName() string {
	return "PayOrderOkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PayOrderOkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayOrderOkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayOrderOkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayOrderOkResponseValidationError{}

// Validate checks the field values on PayNotifyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PayNotifyResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPayNotifies() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PayNotifyResponseValidationError{
					field:  fmt.Sprintf("PayNotifies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PayNotifyResponseValidationError is the validation error returned by
// PayNotifyResponse.Validate if the designated constraints aren't met.
type PayNotifyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayNotifyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayNotifyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayNotifyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayNotifyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayNotifyResponseValidationError) ErrorName() string {
	return "PayNotifyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PayNotifyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayNotifyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayNotifyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayNotifyResponseValidationError{}

// Validate checks the field values on PayNotifyOkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PayNotifyOkResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPayNotifyOks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PayNotifyOkResponseValidationError{
					field:  fmt.Sprintf("PayNotifyOks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PayNotifyOkResponseValidationError is the validation error returned by
// PayNotifyOkResponse.Validate if the designated constraints aren't met.
type PayNotifyOkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayNotifyOkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayNotifyOkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayNotifyOkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayNotifyOkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayNotifyOkResponseValidationError) ErrorName() string {
	return "PayNotifyOkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PayNotifyOkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayNotifyOkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayNotifyOkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayNotifyOkResponseValidationError{}
