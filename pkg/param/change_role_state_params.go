// Copyright (c) ZStack.io, Inc.

package param

// ChangeRoleStateDetailParam ChangeRoleState detail param
type ChangeRoleStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeRoleStateParam ChangeRoleState request param
type ChangeRoleStateParam struct {
	BaseParam
	Params ChangeRoleStateDetailParam `json:"params"`
}
