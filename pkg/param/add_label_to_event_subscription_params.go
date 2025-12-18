// Copyright (c) ZStack.io, Inc.

package param

// AddLabelToEventSubscriptionDetailParam AddLabelToEventSubscription详细参数
type AddLabelToEventSubscriptionDetailParam struct {
	rest string `json:"subscriptionUuid" validate:"required"` // 必填
	rest string `json:"key" validate:"required"` // 必填
	rest string `json:"value" validate:"required"` // 必填
	rest string `json:"operator" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddLabelToEventSubscriptionParam AddLabelToEventSubscription请求参数
type AddLabelToEventSubscriptionParam struct {
	BaseParam
	Params AddLabelToEventSubscriptionDetailParam `json:"params"` // 详细参数
}

