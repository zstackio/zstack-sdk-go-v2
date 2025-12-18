// Copyright (c) ZStack.io, Inc.

package param

// UpdateEventSubscriptionLabelDetailParam UpdateEventSubscriptionLabel detail param
type UpdateEventSubscriptionLabelDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Key string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
	Operator string `json:"operator" validate:"required"`
}

// UpdateEventSubscriptionLabelParam UpdateEventSubscriptionLabel request param
type UpdateEventSubscriptionLabelParam struct {
	BaseParam
	Params UpdateEventSubscriptionLabelDetailParam `json:"params"`
}
