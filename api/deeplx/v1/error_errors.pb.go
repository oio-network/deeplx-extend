// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PallasErrorReason_UNKNOWN.String() && e.Code == 500
}

func ErrorUnknown(format string, args ...interface{}) *errors.Error {
	return errors.New(500, PallasErrorReason_UNKNOWN.String(), fmt.Sprintf(format, args...))
}

func IsInternal(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PallasErrorReason_INTERNAL.String() && e.Code == 500
}

func ErrorInternal(format string, args ...interface{}) *errors.Error {
	return errors.New(500, PallasErrorReason_INTERNAL.String(), fmt.Sprintf(format, args...))
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PallasErrorReason_NOT_FOUND.String() && e.Code == 404
}

func ErrorNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, PallasErrorReason_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PallasErrorReason_CONFLICT.String() && e.Code == 409
}

func ErrorConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(409, PallasErrorReason_CONFLICT.String(), fmt.Sprintf(format, args...))
}