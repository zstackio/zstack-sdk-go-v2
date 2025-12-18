// Copyright (c) ZStack.io, Inc.

package param

// AddVmNicToLoadBalancerDetailParam AddVmNicToLoadBalancer详细参数
type AddVmNicToLoadBalancerDetailParam struct {
	rest []string `json:"vmNicUuids" validate:"required"` // 必填
	rest string `json:"listenerUuid" validate:"required"` // 必填
}

// AddVmNicToLoadBalancerParam AddVmNicToLoadBalancer请求参数
type AddVmNicToLoadBalancerParam struct {
	BaseParam
	Params AddVmNicToLoadBalancerDetailParam `json:"params"` // 详细参数
}

