// Copyright (c) ZStack.io, Inc.

package param

// UpdateHaStrategyConditionDetailParam UpdateHaStrategyCondition detail param
type UpdateHaStrategyConditionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateHaStrategyConditionParam UpdateHaStrategyCondition request param
type UpdateHaStrategyConditionParam struct {
	BaseParam
	Params UpdateHaStrategyConditionDetailParam `json:"params"`
}
