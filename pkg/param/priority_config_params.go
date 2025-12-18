// Copyright (c) ZStack.io, Inc.

package param

// UpdatePriorityConfigDetailParam UpdatePriorityConfig详细参数
type UpdatePriorityConfigDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"cpuShares,omitempty"`
	rest int `json:"oomScoreAdj,omitempty"`
}

// UpdatePriorityConfigParam UpdatePriorityConfig请求参数
type UpdatePriorityConfigParam struct {
	BaseParam
	Params UpdatePriorityConfigDetailParam `json:"params"` // 详细参数
}

