// Copyright (c) ZStack.io, Inc.

package param

// AttachPolicyRouteRuleSetToL3DetailParam AttachPolicyRouteRuleSetToL3 detail param
type AttachPolicyRouteRuleSetToL3DetailParam struct {
	L3Uuid string `json:"l3Uuid" validate:"required"`
	RuleSetUuid string `json:"ruleSetUuid" validate:"required"`
}

// AttachPolicyRouteRuleSetToL3Param AttachPolicyRouteRuleSetToL3 request param
type AttachPolicyRouteRuleSetToL3Param struct {
	BaseParam
	Params AttachPolicyRouteRuleSetToL3DetailParam `json:"params"`
}
