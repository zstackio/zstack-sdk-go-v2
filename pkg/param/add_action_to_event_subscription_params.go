// Copyright (c) ZStack.io, Inc.

package param

// AddActionToEventSubscriptionDetailParam AddActionToEventSubscription detail param
type AddActionToEventSubscriptionDetailParam struct {
	SubscriptionUuid string `json:"subscriptionUuid" validate:"required"`
	ActionUuid string `json:"actionUuid" validate:"required"`
	ActionType string `json:"actionType" validate:"required"`
}

// AddActionToEventSubscriptionParam AddActionToEventSubscription request param
type AddActionToEventSubscriptionParam struct {
	BaseParam
	Params AddActionToEventSubscriptionDetailParam `json:"params"`
}
