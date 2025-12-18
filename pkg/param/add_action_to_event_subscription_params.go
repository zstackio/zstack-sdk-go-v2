// Copyright (c) ZStack.io, Inc.

package param

// AddActionToEventSubscriptionDetailParam AddActionToEventSubscription详细参数
type AddActionToEventSubscriptionDetailParam struct {
	rest string `json:"subscriptionUuid" validate:"required"` // 必填
	rest string `json:"actionUuid" validate:"required"` // 必填
	rest string `json:"actionType" validate:"required"` // 必填
}

// AddActionToEventSubscriptionParam AddActionToEventSubscription请求参数
type AddActionToEventSubscriptionParam struct {
	BaseParam
	Params AddActionToEventSubscriptionDetailParam `json:"params"` // 详细参数
}

