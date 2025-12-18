// Copyright (c) ZStack.io, Inc.

package param

// AddVmToVmSchedulingRuleGroupDetailParam AddVmToVmSchedulingRuleGroup详细参数
type AddVmToVmSchedulingRuleGroupDetailParam struct {
	rest string `json:"vmGroupUuid" validate:"required"` // 必填
	rest string `json:"vmUuid" validate:"required"` // 必填
}

// AddVmToVmSchedulingRuleGroupParam AddVmToVmSchedulingRuleGroup请求参数
type AddVmToVmSchedulingRuleGroupParam struct {
	BaseParam
	Params AddVmToVmSchedulingRuleGroupDetailParam `json:"params"` // 详细参数
}

