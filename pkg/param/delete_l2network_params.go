// Copyright (c) ZStack.io, Inc.

package param

// DeleteL2NetworkDetailParam DeleteL2Network detail param
type DeleteL2NetworkDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteL2NetworkParam DeleteL2Network request param
type DeleteL2NetworkParam struct {
	BaseParam
	Params DeleteL2NetworkDetailParam `json:"params"`
}
