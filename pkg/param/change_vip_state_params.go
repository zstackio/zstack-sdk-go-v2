// Copyright (c) ZStack.io, Inc.

package param

// ChangeVipStateDetailParam ChangeVipState detail param
type ChangeVipStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeVipStateParam ChangeVipState request param
type ChangeVipStateParam struct {
	BaseParam
	Params ChangeVipStateDetailParam `json:"params"`
}
