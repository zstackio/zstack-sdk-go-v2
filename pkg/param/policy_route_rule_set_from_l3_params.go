// Copyright (c) ZStack.io, Inc.

package param

// DetachPolicyRouteRuleSetFromL3DetailParam DetachPolicyRouteRuleSetFromL3详细参数
type DetachPolicyRouteRuleSetFromL3DetailParam struct {
	rest string `json:"l3Uuid" validate:"required"` // 必填
	rest string `json:"ruleSetUuid" validate:"required"` // 必填
}

// DetachPolicyRouteRuleSetFromL3Param DetachPolicyRouteRuleSetFromL3请求参数
type DetachPolicyRouteRuleSetFromL3Param struct {
	BaseParam
	Params DetachPolicyRouteRuleSetFromL3DetailParam `json:"params"` // 详细参数
}

