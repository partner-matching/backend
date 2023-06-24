// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

//  Get_Account_Failed = 1 [(errors.code) = 401];
func IsAddTeamFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PartnerErrorReason_ADD_TEAM_FAILED.String() && e.Code == 500
}

//  Get_Account_Failed = 1 [(errors.code) = 401];
func ErrorAddTeamFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, PartnerErrorReason_ADD_TEAM_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsDeleteTeamFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PartnerErrorReason_DELETE_TEAM_FAILED.String() && e.Code == 500
}

func ErrorDeleteTeamFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, PartnerErrorReason_DELETE_TEAM_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsUpdateTeamFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PartnerErrorReason_UPDATE_TEAM_FAILED.String() && e.Code == 500
}

func ErrorUpdateTeamFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, PartnerErrorReason_UPDATE_TEAM_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetTeamFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PartnerErrorReason_GET_TEAM_FAILED.String() && e.Code == 500
}

func ErrorGetTeamFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, PartnerErrorReason_GET_TEAM_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetTeamListFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PartnerErrorReason_GET_TEAM_LIST_FAILED.String() && e.Code == 500
}

func ErrorGetTeamListFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, PartnerErrorReason_GET_TEAM_LIST_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsAddUserTeamFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PartnerErrorReason_ADD_USER_TEAM_FAILED.String() && e.Code == 500
}

func ErrorAddUserTeamFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, PartnerErrorReason_ADD_USER_TEAM_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsJoinTeamFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == PartnerErrorReason_JOIN_TEAM_FAILED.String() && e.Code == 500
}

func ErrorJoinTeamFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, PartnerErrorReason_JOIN_TEAM_FAILED.String(), fmt.Sprintf(format, args...))
}
