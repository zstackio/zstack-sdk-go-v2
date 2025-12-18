// Copyright (c) ZStack.io, Inc.

package param

// RemoveAccessControlListFromLoadBalancerDetailParam RemoveAccessControlListFromLoadBalancer详细参数
type RemoveAccessControlListFromLoadBalancerDetailParam struct {
	rest []string `json:"aclUuids" validate:"required"` // 必填
	rest string `json:"listenerUuid" validate:"required"` // 必填
	rest []string `json:"serverGroupUuids,omitempty"`
}

// RemoveAccessControlListFromLoadBalancerParam RemoveAccessControlListFromLoadBalancer请求参数
type RemoveAccessControlListFromLoadBalancerParam struct {
	BaseParam
	Params RemoveAccessControlListFromLoadBalancerDetailParam `json:"params"` // 详细参数
}

