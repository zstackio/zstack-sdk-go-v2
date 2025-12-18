// Copyright (c) ZStack.io, Inc.

package param

// DeleteVCenterDetailParam DeleteVCenter detail param
type DeleteVCenterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVCenterParam DeleteVCenter request param
type DeleteVCenterParam struct {
	BaseParam
	Params DeleteVCenterDetailParam `json:"params"`
}
