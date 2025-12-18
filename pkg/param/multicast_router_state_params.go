// Copyright (c) ZStack.io, Inc.

package param

// ChangeMulticastRouterStateDetailParam ChangeMulticastRouterState详细参数
type ChangeMulticastRouterStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeMulticastRouterStateParam ChangeMulticastRouterState请求参数
type ChangeMulticastRouterStateParam struct {
	BaseParam
	Params ChangeMulticastRouterStateDetailParam `json:"params"` // 详细参数
}

