// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelCenterDetailParam DeleteModelCenter detail param
type DeleteModelCenterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteModelCenterParam DeleteModelCenter request param
type DeleteModelCenterParam struct {
	BaseParam
	Params DeleteModelCenterDetailParam `json:"params"`
}
