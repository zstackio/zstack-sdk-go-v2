// Copyright (c) ZStack.io, Inc.

package param

// AddDnsToL3NetworkDetailParam AddDnsToL3Network detail param
type AddDnsToL3NetworkDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Dns string `json:"dns" validate:"required"`
}

// AddDnsToL3NetworkParam AddDnsToL3Network request param
type AddDnsToL3NetworkParam struct {
	BaseParam
	Params AddDnsToL3NetworkDetailParam `json:"params"`
}
