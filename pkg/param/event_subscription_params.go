// Copyright (c) ZStack.io, Inc.

package param

// QueryEventSubscriptionDetailParam QueryEventSubscription详细参数
type QueryEventSubscriptionDetailParam struct {
	rest []interface{} `json:"conditions" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
	rest bool `json:"count,omitempty"`
	rest string `json:"groupBy,omitempty"`
	rest bool `json:"replyWithCount,omitempty"`
	rest string `json:"filterName,omitempty"`
	rest string `json:"sortBy,omitempty"`
	rest string `json:"sortDirection,omitempty"`
	rest []string `json:"fields,omitempty"`
}

// QueryEventSubscriptionParam QueryEventSubscription请求参数
type QueryEventSubscriptionParam struct {
	BaseParam
	Params QueryEventSubscriptionDetailParam `json:"params"` // 详细参数
}

