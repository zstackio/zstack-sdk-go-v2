// Copyright (c) ZStack.io, Inc.

package param

// RemoveLabelFromEventSubscriptionDetailParam RemoveLabelFromEventSubscription详细参数
type RemoveLabelFromEventSubscriptionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// RemoveLabelFromEventSubscriptionParam RemoveLabelFromEventSubscription请求参数
type RemoveLabelFromEventSubscriptionParam struct {
	BaseParam
	Params RemoveLabelFromEventSubscriptionDetailParam `json:"params"` // 详细参数
}

