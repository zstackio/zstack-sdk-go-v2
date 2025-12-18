// Copyright (c) ZStack.io, Inc.

package param

// AddVmNicToLoadBalancerDetailParam AddVmNicToLoadBalancer detail param
type AddVmNicToLoadBalancerDetailParam struct {
	VmNicUuids []string `json:"vmNicUuids" validate:"required"`
	ListenerUuid string `json:"listenerUuid" validate:"required"`
}

// AddVmNicToLoadBalancerParam AddVmNicToLoadBalancer request param
type AddVmNicToLoadBalancerParam struct {
	BaseParam
	Params AddVmNicToLoadBalancerDetailParam `json:"params"`
}
