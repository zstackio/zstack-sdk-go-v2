// Copyright (c) ZStack.io, Inc.

package param

// GetVmsSchedulingStateFromSchedulingRuleDetailParam GetVmsSchedulingStateFromSchedulingRule详细参数
type GetVmsSchedulingStateFromSchedulingRuleDetailParam struct {
	rest string `json:"ruleUuid" validate:"required"` // 必填
	rest []string `json:"vmUuids" validate:"required"` // 必填
}

// GetVmsSchedulingStateFromSchedulingRuleParam GetVmsSchedulingStateFromSchedulingRule请求参数
type GetVmsSchedulingStateFromSchedulingRuleParam struct {
	BaseParam
	Params GetVmsSchedulingStateFromSchedulingRuleDetailParam `json:"params"` // 详细参数
}

