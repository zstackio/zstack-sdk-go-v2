// Copyright (c) ZStack.io, Inc.

package param

// DeleteIAM2VirtualIDGroupDetailParam DeleteIAM2VirtualIDGroup detail param
type DeleteIAM2VirtualIDGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteIAM2VirtualIDGroupParam DeleteIAM2VirtualIDGroup request param
type DeleteIAM2VirtualIDGroupParam struct {
	BaseParam
	Params DeleteIAM2VirtualIDGroupDetailParam `json:"params"`
}
