// Copyright (c) ZStack.io, Inc.

package param

// DeleteHostSchedulingRuleGroupDetailParam DeleteHostSchedulingRuleGroup detail param
type DeleteHostSchedulingRuleGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteHostSchedulingRuleGroupParam DeleteHostSchedulingRuleGroup request param
type DeleteHostSchedulingRuleGroupParam struct {
	BaseParam
	Params DeleteHostSchedulingRuleGroupDetailParam `json:"params"`
}
