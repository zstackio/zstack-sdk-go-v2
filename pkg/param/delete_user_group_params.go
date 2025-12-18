// Copyright (c) ZStack.io, Inc.

package param

// DeleteUserGroupDetailParam DeleteUserGroup detail param
type DeleteUserGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteUserGroupParam DeleteUserGroup request param
type DeleteUserGroupParam struct {
	BaseParam
	Params DeleteUserGroupDetailParam `json:"params"`
}
