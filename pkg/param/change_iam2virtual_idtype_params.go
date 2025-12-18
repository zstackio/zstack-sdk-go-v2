// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2VirtualIDTypeDetailParam ChangeIAM2VirtualIDType detail param
type ChangeIAM2VirtualIDTypeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Type string `json:"type" validate:"required"`
}

// ChangeIAM2VirtualIDTypeParam ChangeIAM2VirtualIDType request param
type ChangeIAM2VirtualIDTypeParam struct {
	BaseParam
	Params ChangeIAM2VirtualIDTypeDetailParam `json:"params"`
}
