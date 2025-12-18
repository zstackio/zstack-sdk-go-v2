// Copyright (c) ZStack.io, Inc.

package param

// UpdateEventSubscriptionLabelDetailParam UpdateEventSubscriptionLabel详细参数
type UpdateEventSubscriptionLabelDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"key" validate:"required"` // 必填
	rest string `json:"value" validate:"required"` // 必填
	rest string `json:"operator" validate:"required"` // 必填
}

// UpdateEventSubscriptionLabelParam UpdateEventSubscriptionLabel请求参数
type UpdateEventSubscriptionLabelParam struct {
	BaseParam
	Params UpdateEventSubscriptionLabelDetailParam `json:"params"` // 详细参数
}

