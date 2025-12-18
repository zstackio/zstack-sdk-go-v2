// Copyright (c) ZStack.io, Inc.

package param

// ChangeSecurityGroupStateDetailParam ChangeSecurityGroupState detail param
type ChangeSecurityGroupStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSecurityGroupStateParam ChangeSecurityGroupState request param
type ChangeSecurityGroupStateParam struct {
	BaseParam
	Params ChangeSecurityGroupStateDetailParam `json:"params"`
}
