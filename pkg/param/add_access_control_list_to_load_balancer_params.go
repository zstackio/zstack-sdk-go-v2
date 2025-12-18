// Copyright (c) ZStack.io, Inc.

package param

// AddAccessControlListToLoadBalancerDetailParam AddAccessControlListToLoadBalancer详细参数
type AddAccessControlListToLoadBalancerDetailParam struct {
	rest []string `json:"aclUuids" validate:"required"` // 必填
	rest string `json:"aclType" validate:"required"` // 必填
	rest string `json:"listenerUuid" validate:"required"` // 必填
	rest []string `json:"serverGroupUuids,omitempty"`
}

// AddAccessControlListToLoadBalancerParam AddAccessControlListToLoadBalancer请求参数
type AddAccessControlListToLoadBalancerParam struct {
	BaseParam
	Params AddAccessControlListToLoadBalancerDetailParam `json:"params"` // 详细参数
}

