// Copyright (c) ZStack.io, Inc.

package param

// ChangeEventSubscriptionStateDetailParam ChangeEventSubscriptionState detail param
type ChangeEventSubscriptionStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	State string `json:"state" validate:"required"`
}

// ChangeEventSubscriptionStateParam ChangeEventSubscriptionState request param
type ChangeEventSubscriptionStateParam struct {
	BaseParam
	Params ChangeEventSubscriptionStateDetailParam `json:"params"`
}
