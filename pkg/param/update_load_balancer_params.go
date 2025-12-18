// Copyright (c) ZStack.io, Inc.

package param

// UpdateLoadBalancerDetailParam UpdateLoadBalancer detail param
type UpdateLoadBalancerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateLoadBalancerParam UpdateLoadBalancer request param
type UpdateLoadBalancerParam struct {
	BaseParam
	Params UpdateLoadBalancerDetailParam `json:"params"`
}
