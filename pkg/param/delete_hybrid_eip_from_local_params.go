// Copyright (c) ZStack.io, Inc.

package param

// DeleteHybridEipFromLocalDetailParam DeleteHybridEipFromLocal detail param
type DeleteHybridEipFromLocalDetailParam struct {
	Type string `json:"type" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHybridEipFromLocalParam DeleteHybridEipFromLocal request param
type DeleteHybridEipFromLocalParam struct {
	BaseParam
	Params DeleteHybridEipFromLocalDetailParam `json:"params"`
}
