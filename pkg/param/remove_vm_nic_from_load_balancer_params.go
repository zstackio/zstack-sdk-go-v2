// Copyright (c) ZStack.io, Inc.

package param

// RemoveVmNicFromLoadBalancerDetailParam RemoveVmNicFromLoadBalancer detail param
type RemoveVmNicFromLoadBalancerDetailParam struct {
	VmNicUuids []string `json:"vmNicUuids" validate:"required"`
	ListenerUuid string `json:"listenerUuid" validate:"required"`
}

// RemoveVmNicFromLoadBalancerParam RemoveVmNicFromLoadBalancer request param
type RemoveVmNicFromLoadBalancerParam struct {
	BaseParam
	Params RemoveVmNicFromLoadBalancerDetailParam `json:"params"`
}
