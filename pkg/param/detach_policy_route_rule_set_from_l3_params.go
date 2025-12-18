// Copyright (c) ZStack.io, Inc.

package param

// DetachPolicyRouteRuleSetFromL3DetailParam DetachPolicyRouteRuleSetFromL3 detail param
type DetachPolicyRouteRuleSetFromL3DetailParam struct {
	L3Uuid string `json:"l3Uuid" validate:"required"`
	RuleSetUuid string `json:"ruleSetUuid" validate:"required"`
}

// DetachPolicyRouteRuleSetFromL3Param DetachPolicyRouteRuleSetFromL3 request param
type DetachPolicyRouteRuleSetFromL3Param struct {
	BaseParam
	Params DetachPolicyRouteRuleSetFromL3DetailParam `json:"params"`
}
