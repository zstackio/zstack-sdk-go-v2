// Copyright (c) ZStack.io, Inc.

package param

// GetAvailableVpcL3NetworkDetailParam GetAvailableVpcL3Network detail param
type GetAvailableVpcL3NetworkDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
}

// GetAvailableVpcL3NetworkParam GetAvailableVpcL3Network request param
type GetAvailableVpcL3NetworkParam struct {
	BaseParam
	Params GetAvailableVpcL3NetworkDetailParam `json:"params"`
}
