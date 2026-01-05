// Copyright (c) ZStack.io, Inc.

package param

// ReconnectNfvInstDetailParam ReconnectNfvInst detail param
type ReconnectNfvInstDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// ReconnectNfvInstParam ReconnectNfvInst request param
type ReconnectNfvInstParam struct {
	BaseParam
	Params ReconnectNfvInstDetailParam `json:"params"`
}
