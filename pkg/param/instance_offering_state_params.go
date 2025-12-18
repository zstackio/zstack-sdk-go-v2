// Copyright (c) ZStack.io, Inc.

package param

// ChangeInstanceOfferingStateDetailParam ChangeInstanceOfferingState详细参数
type ChangeInstanceOfferingStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeInstanceOfferingStateParam ChangeInstanceOfferingState请求参数
type ChangeInstanceOfferingStateParam struct {
	BaseParam
	Params ChangeInstanceOfferingStateDetailParam `json:"params"` // 详细参数
}

