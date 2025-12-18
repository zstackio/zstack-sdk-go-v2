// Copyright (c) ZStack.io, Inc.

package param

// GetVmsSchedulingStateFromSchedulingRuleDetailParam GetVmsSchedulingStateFromSchedulingRule detail param
type GetVmsSchedulingStateFromSchedulingRuleDetailParam struct {
	RuleUuid string `json:"ruleUuid" validate:"required"`
	VmUuids []string `json:"vmUuids" validate:"required"`
}

// GetVmsSchedulingStateFromSchedulingRuleParam GetVmsSchedulingStateFromSchedulingRule request param
type GetVmsSchedulingStateFromSchedulingRuleParam struct {
	BaseParam
	Params GetVmsSchedulingStateFromSchedulingRuleDetailParam `json:"params"`
}
