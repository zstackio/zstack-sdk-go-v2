// Copyright (c) ZStack.io, Inc.

package param

// AddDnsToL3NetworkDetailParam AddDnsToL3Network详细参数
type AddDnsToL3NetworkDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"dns" validate:"required"` // 必填
}

// AddDnsToL3NetworkParam AddDnsToL3Network请求参数
type AddDnsToL3NetworkParam struct {
	BaseParam
	Params AddDnsToL3NetworkDetailParam `json:"params"` // 详细参数
}

