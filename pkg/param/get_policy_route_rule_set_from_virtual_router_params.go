// Copyright (c) ZStack.io, Inc.

package param

// GetPolicyRouteRuleSetFromVirtualRouterDetailParam GetPolicyRouteRuleSetFromVirtualRouter detail param
type GetPolicyRouteRuleSetFromVirtualRouterDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// GetPolicyRouteRuleSetFromVirtualRouterParam GetPolicyRouteRuleSetFromVirtualRouter request param
type GetPolicyRouteRuleSetFromVirtualRouterParam struct {
	BaseParam
	Params GetPolicyRouteRuleSetFromVirtualRouterDetailParam `json:"params"`
}
