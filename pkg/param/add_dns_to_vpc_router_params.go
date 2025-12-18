// Copyright (c) ZStack.io, Inc.

package param

// AddDnsToVpcRouterDetailParam AddDnsToVpcRouter详细参数
type AddDnsToVpcRouterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"dns" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddDnsToVpcRouterParam AddDnsToVpcRouter请求参数
type AddDnsToVpcRouterParam struct {
	BaseParam
	Params AddDnsToVpcRouterDetailParam `json:"params"` // 详细参数
}

