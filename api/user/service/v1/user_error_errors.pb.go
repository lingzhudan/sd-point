// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 用户名错误
func IsAccountError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ACCOUNT_ERROR.String() && e.Code == 401
}

// 用户名错误
func ErrorAccountError(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_ACCOUNT_ERROR.String(), fmt.Sprintf(format, args...))
}

// 密码错误
func IsPasswordError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PASSWORD_ERROR.String() && e.Code == 401
}

// 密码错误
func ErrorPasswordError(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_PASSWORD_ERROR.String(), fmt.Sprintf(format, args...))
}

// 无此sessionID
func IsSessionIdNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SESSION_ID_NOT_FOUND.String() && e.Code == 404
}

// 无此sessionID
func ErrorSessionIdNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_SESSION_ID_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 账号已注册
func IsAccountRegistered(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ACCOUNT_REGISTERED.String() && e.Code == 401
}

// 账号已注册
func ErrorAccountRegistered(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_ACCOUNT_REGISTERED.String(), fmt.Sprintf(format, args...))
}

// 非法账号名
func IsAccountInvalid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ACCOUNT_INVALID.String() && e.Code == 401
}

// 非法账号名
func ErrorAccountInvalid(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_ACCOUNT_INVALID.String(), fmt.Sprintf(format, args...))
}

// 非法密码
func IsPasswordInvalid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PASSWORD_INVALID.String() && e.Code == 401
}

// 非法密码
func ErrorPasswordInvalid(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_PASSWORD_INVALID.String(), fmt.Sprintf(format, args...))
}

// 微信已注册
func IsWechatRegistered(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_WECHAT_REGISTERED.String() && e.Code == 401
}

// 微信已注册
func ErrorWechatRegistered(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_WECHAT_REGISTERED.String(), fmt.Sprintf(format, args...))
}

// 微信code错误
func IsWechatCodeError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_WECHAT_CODE_ERROR.String() && e.Code == 401
}

// 微信code错误
func ErrorWechatCodeError(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_WECHAT_CODE_ERROR.String(), fmt.Sprintf(format, args...))
}

// 无此用户
func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

// 无此用户
func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 微信已注册
func IsPhoneNumberRegistered(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PHONE_NUMBER_REGISTERED.String() && e.Code == 401
}

// 微信已注册
func ErrorPhoneNumberRegistered(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_PHONE_NUMBER_REGISTERED.String(), fmt.Sprintf(format, args...))
}