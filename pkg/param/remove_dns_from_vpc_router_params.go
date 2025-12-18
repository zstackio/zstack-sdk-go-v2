// Copyright (c) ZStack.io, Inc.

package param

// RemoveDnsFromVpcRouterDetailParam RemoveDnsFromVpcRouter detail param
type RemoveDnsFromVpcRouterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Dns string `json:"dns" validate:"required"`
}

// RemoveDnsFromVpcRouterParam RemoveDnsFromVpcRouter request param
type RemoveDnsFromVpcRouterParam struct {
	BaseParam
	Params RemoveDnsFromVpcRouterDetailParam `json:"params"`
}
