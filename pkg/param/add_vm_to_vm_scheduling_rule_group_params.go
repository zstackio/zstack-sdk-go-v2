// Copyright (c) ZStack.io, Inc.

package param

// AddVmToVmSchedulingRuleGroupDetailParam AddVmToVmSchedulingRuleGroup detail param
type AddVmToVmSchedulingRuleGroupDetailParam struct {
	VmGroupUuid string `json:"vmGroupUuid" validate:"required"`
	VmUuid string `json:"vmUuid" validate:"required"`
}

// AddVmToVmSchedulingRuleGroupParam AddVmToVmSchedulingRuleGroup request param
type AddVmToVmSchedulingRuleGroupParam struct {
	BaseParam
	Params AddVmToVmSchedulingRuleGroupDetailParam `json:"params"`
}
