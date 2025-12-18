// Copyright (c) ZStack.io, Inc.

package param

// QueryModelEvalServiceInstanceGroupDetailParam QueryModelEvalServiceInstanceGroup详细参数
type QueryModelEvalServiceInstanceGroupDetailParam struct {
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

// QueryModelEvalServiceInstanceGroupParam QueryModelEvalServiceInstanceGroup请求参数
type QueryModelEvalServiceInstanceGroupParam struct {
	BaseParam
	Params QueryModelEvalServiceInstanceGroupDetailParam `json:"params"` // 详细参数
}

