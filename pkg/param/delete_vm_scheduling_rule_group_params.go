// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmSchedulingRuleGroupDetailParam DeleteVmSchedulingRuleGroup detail param
type DeleteVmSchedulingRuleGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteVmSchedulingRuleGroupParam DeleteVmSchedulingRuleGroup request param
type DeleteVmSchedulingRuleGroupParam struct {
	BaseParam
	Params DeleteVmSchedulingRuleGroupDetailParam `json:"params"`
}
