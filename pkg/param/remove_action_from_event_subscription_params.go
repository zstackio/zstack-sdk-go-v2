// Copyright (c) ZStack.io, Inc.

package param

// RemoveActionFromEventSubscriptionDetailParam RemoveActionFromEventSubscription detail param
type RemoveActionFromEventSubscriptionDetailParam struct {
	SubscriptionUuid string `json:"subscriptionUuid" validate:"required"`
	ActionUuid string `json:"actionUuid" validate:"required"`
}

// RemoveActionFromEventSubscriptionParam RemoveActionFromEventSubscription request param
type RemoveActionFromEventSubscriptionParam struct {
	BaseParam
	Params RemoveActionFromEventSubscriptionDetailParam `json:"params"`
}
