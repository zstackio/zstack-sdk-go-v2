// Copyright (c) ZStack.io, Inc.

package param

// AddLabelToEventSubscriptionDetailParam AddLabelToEventSubscription detail param
type AddLabelToEventSubscriptionDetailParam struct {
	SubscriptionUuid string `json:"subscriptionUuid" validate:"required"`
	Key string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
	Operator string `json:"operator" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddLabelToEventSubscriptionParam AddLabelToEventSubscription request param
type AddLabelToEventSubscriptionParam struct {
	BaseParam
	Params AddLabelToEventSubscriptionDetailParam `json:"params"`
}
