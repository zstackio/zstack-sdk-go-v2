// Copyright (c) ZStack.io, Inc.

package param

// DetachVmFromVmSchedulingRuleGroupDetailParam DetachVmFromVmSchedulingRuleGroup detail param
type DetachVmFromVmSchedulingRuleGroupDetailParam struct {
	VmGroupUuid string `json:"vmGroupUuid" validate:"required"`
	VmUuid string `json:"vmUuid" validate:"required"`
}

// DetachVmFromVmSchedulingRuleGroupParam DetachVmFromVmSchedulingRuleGroup request param
type DetachVmFromVmSchedulingRuleGroupParam struct {
	BaseParam
	Params DetachVmFromVmSchedulingRuleGroupDetailParam `json:"params"`
}
