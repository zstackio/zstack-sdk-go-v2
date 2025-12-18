// Copyright (c) ZStack.io, Inc.

package param

// DeleteVxlanL2NetworkDetailParam DeleteVxlanL2Network detail param
type DeleteVxlanL2NetworkDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVxlanL2NetworkParam DeleteVxlanL2Network request param
type DeleteVxlanL2NetworkParam struct {
	BaseParam
	Params DeleteVxlanL2NetworkDetailParam `json:"params"`
}
