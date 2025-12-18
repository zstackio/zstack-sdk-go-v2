// Copyright (c) ZStack.io, Inc.

package param

// UpdateHaStrategyConditionDetailParam UpdateHaStrategyCondition详细参数
type UpdateHaStrategyConditionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"state,omitempty"`
}

// UpdateHaStrategyConditionParam UpdateHaStrategyCondition请求参数
type UpdateHaStrategyConditionParam struct {
	BaseParam
	Params UpdateHaStrategyConditionDetailParam `json:"params"` // 详细参数
}

