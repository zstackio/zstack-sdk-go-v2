// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2ChassisOfferingStateDetailParam ChangeBareMetal2ChassisOfferingState详细参数
type ChangeBareMetal2ChassisOfferingStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeBareMetal2ChassisOfferingStateParam ChangeBareMetal2ChassisOfferingState请求参数
type ChangeBareMetal2ChassisOfferingStateParam struct {
	BaseParam
	Params ChangeBareMetal2ChassisOfferingStateDetailParam `json:"params"` // 详细参数
}

