// Copyright (c) ZStack.io, Inc.

package param

// RemoveVmNicFromLoadBalancerDetailParam RemoveVmNicFromLoadBalancer详细参数
type RemoveVmNicFromLoadBalancerDetailParam struct {
	rest []string `json:"vmNicUuids" validate:"required"` // 必填
	rest string `json:"listenerUuid" validate:"required"` // 必填
}

// RemoveVmNicFromLoadBalancerParam RemoveVmNicFromLoadBalancer请求参数
type RemoveVmNicFromLoadBalancerParam struct {
	BaseParam
	Params RemoveVmNicFromLoadBalancerDetailParam `json:"params"` // 详细参数
}

