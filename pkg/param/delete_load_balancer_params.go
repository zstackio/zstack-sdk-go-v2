// Copyright (c) ZStack.io, Inc.

package param

// DeleteLoadBalancerDetailParam DeleteLoadBalancer detail param
type DeleteLoadBalancerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteLoadBalancerParam DeleteLoadBalancer request param
type DeleteLoadBalancerParam struct {
	BaseParam
	Params DeleteLoadBalancerDetailParam `json:"params"`
}
