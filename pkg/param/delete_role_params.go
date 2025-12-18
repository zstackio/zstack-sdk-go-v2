// Copyright (c) ZStack.io, Inc.

package param

// DeleteRoleDetailParam DeleteRole detail param
type DeleteRoleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteRoleParam DeleteRole request param
type DeleteRoleParam struct {
	BaseParam
	Params DeleteRoleDetailParam `json:"params"`
}
