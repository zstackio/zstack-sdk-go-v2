// Copyright (c) ZStack.io, Inc.

package param

// ValidateVmSchedulingRuleDetailParam ValidateVmSchedulingRule详细参数
type ValidateVmSchedulingRuleDetailParam struct {
	rest string `json:"vmGroupUuid" validate:"required"` // 必填
	rest string `json:"hostGroupUuid,omitempty"`
	rest string `json:"rule" validate:"required"` // 必填
	rest string `json:"mode" validate:"required"` // 必填
	rest string `json:"zoneUuid,omitempty"`
}

// ValidateVmSchedulingRuleParam ValidateVmSchedulingRule请求参数
type ValidateVmSchedulingRuleParam struct {
	BaseParam
	Params ValidateVmSchedulingRuleDetailParam `json:"params"` // 详细参数
}

