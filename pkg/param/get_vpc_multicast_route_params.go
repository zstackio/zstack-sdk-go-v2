// Copyright (c) ZStack.io, Inc.

package param

// GetVpcMulticastRouteDetailParam GetVpcMulticastRoute detail param
type GetVpcMulticastRouteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVpcMulticastRouteParam GetVpcMulticastRoute request param
type GetVpcMulticastRouteParam struct {
	BaseParam
	Params GetVpcMulticastRouteDetailParam `json:"params"`
}
