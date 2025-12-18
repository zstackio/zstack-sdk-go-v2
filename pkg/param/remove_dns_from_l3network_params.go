// Copyright (c) ZStack.io, Inc.

package param

// RemoveDnsFromL3NetworkDetailParam RemoveDnsFromL3Network detail param
type RemoveDnsFromL3NetworkDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Dns string `json:"dns" validate:"required"`
}

// RemoveDnsFromL3NetworkParam RemoveDnsFromL3Network request param
type RemoveDnsFromL3NetworkParam struct {
	BaseParam
	Params RemoveDnsFromL3NetworkDetailParam `json:"params"`
}
