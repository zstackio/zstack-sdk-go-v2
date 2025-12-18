// Copyright (c) ZStack.io, Inc.

package param

// DetachVmFromVmSchedulingRuleGroupDetailParam DetachVmFromVmSchedulingRuleGroup详细参数
type DetachVmFromVmSchedulingRuleGroupDetailParam struct {
	rest string `json:"vmGroupUuid" validate:"required"` // 必填
	rest string `json:"vmUuid" validate:"required"` // 必填
}

// DetachVmFromVmSchedulingRuleGroupParam DetachVmFromVmSchedulingRuleGroup请求参数
type DetachVmFromVmSchedulingRuleGroupParam struct {
	BaseParam
	Params DetachVmFromVmSchedulingRuleGroupDetailParam `json:"params"` // 详细参数
}

