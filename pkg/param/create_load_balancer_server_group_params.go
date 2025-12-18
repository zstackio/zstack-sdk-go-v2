// Copyright (c) ZStack.io, Inc.

package param

// CreateLoadBalancerServerGroupDetailParam CreateLoadBalancerServerGroup detail param
type CreateLoadBalancerServerGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	LoadBalancerUuid string `json:"loadBalancerUuid" validate:"required"`
	IpVersion int `json:"ipVersion,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateLoadBalancerServerGroupParam CreateLoadBalancerServerGroup request param
type CreateLoadBalancerServerGroupParam struct {
	BaseParam
	Params CreateLoadBalancerServerGroupDetailParam `json:"params"`
}
