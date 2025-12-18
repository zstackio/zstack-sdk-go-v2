// Copyright (c) ZStack.io, Inc.

package param

// ChangeEventSubscriptionStateDetailParam ChangeEventSubscriptionState详细参数
type ChangeEventSubscriptionStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"state" validate:"required"` // 必填
}

// ChangeEventSubscriptionStateParam ChangeEventSubscriptionState请求参数
type ChangeEventSubscriptionStateParam struct {
	BaseParam
	Params ChangeEventSubscriptionStateDetailParam `json:"params"` // 详细参数
}

