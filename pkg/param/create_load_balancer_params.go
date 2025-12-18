// Copyright (c) ZStack.io, Inc.

package param

// CreateLoadBalancerDetailParam CreateLoadBalancer detail param
type CreateLoadBalancerDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	Ipv6VipUuid string `json:"ipv6VipUuid,omitempty"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateLoadBalancerParam CreateLoadBalancer request param
type CreateLoadBalancerParam struct {
	BaseParam
	Params CreateLoadBalancerDetailParam `json:"params"`
}
