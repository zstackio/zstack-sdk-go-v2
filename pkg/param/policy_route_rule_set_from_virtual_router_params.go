// Copyright (c) ZStack.io, Inc.

package param

// GetPolicyRouteRuleSetFromVirtualRouterDetailParam GetPolicyRouteRuleSetFromVirtualRouter详细参数
type GetPolicyRouteRuleSetFromVirtualRouterDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// GetPolicyRouteRuleSetFromVirtualRouterParam GetPolicyRouteRuleSetFromVirtualRouter请求参数
type GetPolicyRouteRuleSetFromVirtualRouterParam struct {
	BaseParam
	Params GetPolicyRouteRuleSetFromVirtualRouterDetailParam `json:"params"` // 详细参数
}

