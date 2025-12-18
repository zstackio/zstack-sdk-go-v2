// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmSchedulingRuleGroupDetailParam DeleteVmSchedulingRuleGroup详细参数
type DeleteVmSchedulingRuleGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// DeleteVmSchedulingRuleGroupParam DeleteVmSchedulingRuleGroup请求参数
type DeleteVmSchedulingRuleGroupParam struct {
	BaseParam
	Params DeleteVmSchedulingRuleGroupDetailParam `json:"params"` // 详细参数
}

