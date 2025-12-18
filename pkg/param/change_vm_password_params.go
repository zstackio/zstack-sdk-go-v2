// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmPasswordDetailParam ChangeVmPassword detail param
type ChangeVmPasswordDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Password string `json:"password" validate:"required"`
	Account string `json:"account" validate:"required"`
}

// ChangeVmPasswordParam ChangeVmPassword request param
type ChangeVmPasswordParam struct {
	BaseParam
	Params ChangeVmPasswordDetailParam `json:"params"`
}
