// Copyright (c) ZStack.io, Inc.

package param

// RemoveVmSchedulingRuleDetailParam RemoveVmSchedulingRule详细参数
type RemoveVmSchedulingRuleDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveVmSchedulingRuleParam RemoveVmSchedulingRule请求参数
type RemoveVmSchedulingRuleParam struct {
	BaseParam
	Params RemoveVmSchedulingRuleDetailParam `json:"params"` // 详细参数
}

