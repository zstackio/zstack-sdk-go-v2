// Copyright (c) ZStack.io, Inc.

package param

// ChangeLoadBalancerBackendServerDetailParam ChangeLoadBalancerBackendServer detail param
type ChangeLoadBalancerBackendServerDetailParam struct {
	ServerGroupUuid string `json:"serverGroupUuid" validate:"required"`
	VmNics []interface{} `json:"vmNics,omitempty"`
	Servers []interface{} `json:"servers,omitempty"`
}

// ChangeLoadBalancerBackendServerParam ChangeLoadBalancerBackendServer request param
type ChangeLoadBalancerBackendServerParam struct {
	BaseParam
	Params ChangeLoadBalancerBackendServerDetailParam `json:"params"`
}
