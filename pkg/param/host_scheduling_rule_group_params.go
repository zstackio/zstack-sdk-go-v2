// Copyright (c) ZStack.io, Inc.

package param

// DeleteHostSchedulingRuleGroupDetailParam DeleteHostSchedulingRuleGroup详细参数
type DeleteHostSchedulingRuleGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// DeleteHostSchedulingRuleGroupParam DeleteHostSchedulingRuleGroup请求参数
type DeleteHostSchedulingRuleGroupParam struct {
	BaseParam
	Params DeleteHostSchedulingRuleGroupDetailParam `json:"params"` // 详细参数
}

