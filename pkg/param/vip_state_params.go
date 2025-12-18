// Copyright (c) ZStack.io, Inc.

package param

// ChangeVipStateDetailParam ChangeVipState详细参数
type ChangeVipStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeVipStateParam ChangeVipState请求参数
type ChangeVipStateParam struct {
	BaseParam
	Params ChangeVipStateDetailParam `json:"params"` // 详细参数
}

