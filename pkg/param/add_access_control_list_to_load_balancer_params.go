// Copyright (c) ZStack.io, Inc.

package param

// AddAccessControlListToLoadBalancerDetailParam AddAccessControlListToLoadBalancer detail param
type AddAccessControlListToLoadBalancerDetailParam struct {
	AclUuids []string `json:"aclUuids" validate:"required"`
	AclType string `json:"aclType" validate:"required"`
	ListenerUuid string `json:"listenerUuid" validate:"required"`
	ServerGroupUuids []string `json:"serverGroupUuids,omitempty"`
}

// AddAccessControlListToLoadBalancerParam AddAccessControlListToLoadBalancer request param
type AddAccessControlListToLoadBalancerParam struct {
	BaseParam
	Params AddAccessControlListToLoadBalancerDetailParam `json:"params"`
}
