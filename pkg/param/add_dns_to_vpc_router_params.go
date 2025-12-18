// Copyright (c) ZStack.io, Inc.

package param

// AddDnsToVpcRouterDetailParam AddDnsToVpcRouter detail param
type AddDnsToVpcRouterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Dns string `json:"dns" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddDnsToVpcRouterParam AddDnsToVpcRouter request param
type AddDnsToVpcRouterParam struct {
	BaseParam
	Params AddDnsToVpcRouterDetailParam `json:"params"`
}
