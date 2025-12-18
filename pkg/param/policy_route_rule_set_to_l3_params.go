// Copyright (c) ZStack.io, Inc.

package param

// AttachPolicyRouteRuleSetToL3DetailParam AttachPolicyRouteRuleSetToL3详细参数
type AttachPolicyRouteRuleSetToL3DetailParam struct {
	rest string `json:"l3Uuid" validate:"required"` // 必填
	rest string `json:"ruleSetUuid" validate:"required"` // 必填
}

// AttachPolicyRouteRuleSetToL3Param AttachPolicyRouteRuleSetToL3请求参数
type AttachPolicyRouteRuleSetToL3Param struct {
	BaseParam
	Params AttachPolicyRouteRuleSetToL3DetailParam `json:"params"` // 详细参数
}

