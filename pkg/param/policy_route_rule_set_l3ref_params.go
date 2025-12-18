// Copyright (c) ZStack.io, Inc.

package param

// QueryPolicyRouteRuleSetL3RefDetailParam QueryPolicyRouteRuleSetL3Ref详细参数
type QueryPolicyRouteRuleSetL3RefDetailParam struct {
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

// QueryPolicyRouteRuleSetL3RefParam QueryPolicyRouteRuleSetL3Ref请求参数
type QueryPolicyRouteRuleSetL3RefParam struct {
	BaseParam
	Params QueryPolicyRouteRuleSetL3RefDetailParam `json:"params"` // 详细参数
}

