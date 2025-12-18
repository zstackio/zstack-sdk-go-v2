// Copyright (c) ZStack.io, Inc.

package param

// RemoveAccessControlListFromLoadBalancerDetailParam RemoveAccessControlListFromLoadBalancer detail param
type RemoveAccessControlListFromLoadBalancerDetailParam struct {
	AclUuids []string `json:"aclUuids" validate:"required"`
	ListenerUuid string `json:"listenerUuid" validate:"required"`
	ServerGroupUuids []string `json:"serverGroupUuids,omitempty"`
}

// RemoveAccessControlListFromLoadBalancerParam RemoveAccessControlListFromLoadBalancer request param
type RemoveAccessControlListFromLoadBalancerParam struct {
	BaseParam
	Params RemoveAccessControlListFromLoadBalancerDetailParam `json:"params"`
}
