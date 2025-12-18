// Copyright (c) ZStack.io, Inc.

package param

// DeleteIAM2VirtualIDDetailParam DeleteIAM2VirtualID detail param
type DeleteIAM2VirtualIDDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteIAM2VirtualIDParam DeleteIAM2VirtualID request param
type DeleteIAM2VirtualIDParam struct {
	BaseParam
	Params DeleteIAM2VirtualIDDetailParam `json:"params"`
}
