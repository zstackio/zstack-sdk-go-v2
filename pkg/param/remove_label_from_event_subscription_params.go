// Copyright (c) ZStack.io, Inc.

package param

// RemoveLabelFromEventSubscriptionDetailParam RemoveLabelFromEventSubscription detail param
type RemoveLabelFromEventSubscriptionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RemoveLabelFromEventSubscriptionParam RemoveLabelFromEventSubscription request param
type RemoveLabelFromEventSubscriptionParam struct {
	BaseParam
	Params RemoveLabelFromEventSubscriptionDetailParam `json:"params"`
}
