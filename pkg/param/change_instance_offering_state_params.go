// Copyright (c) ZStack.io, Inc.

package param

// ChangeInstanceOfferingStateDetailParam ChangeInstanceOfferingState detail param
type ChangeInstanceOfferingStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeInstanceOfferingStateParam ChangeInstanceOfferingState request param
type ChangeInstanceOfferingStateParam struct {
	BaseParam
	Params ChangeInstanceOfferingStateDetailParam `json:"params"`
}
