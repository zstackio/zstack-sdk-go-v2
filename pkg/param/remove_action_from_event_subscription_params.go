// Copyright (c) ZStack.io, Inc.

package param

// RemoveActionFromEventSubscriptionDetailParam RemoveActionFromEventSubscription详细参数
type RemoveActionFromEventSubscriptionDetailParam struct {
	rest string `json:"subscriptionUuid" validate:"required"` // 必填
	rest string `json:"actionUuid" validate:"required"` // 必填
}

// RemoveActionFromEventSubscriptionParam RemoveActionFromEventSubscription请求参数
type RemoveActionFromEventSubscriptionParam struct {
	BaseParam
	Params RemoveActionFromEventSubscriptionDetailParam `json:"params"` // 详细参数
}

