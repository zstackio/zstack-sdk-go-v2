// Copyright (c) ZStack.io, Inc.

package param

// DeleteUserDetailParam DeleteUser detail param
type DeleteUserDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteUserParam DeleteUser request param
type DeleteUserParam struct {
	BaseParam
	Params DeleteUserDetailParam `json:"params"`
}
