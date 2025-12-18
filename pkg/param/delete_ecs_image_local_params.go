// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsImageLocalDetailParam DeleteEcsImageLocal detail param
type DeleteEcsImageLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsImageLocalParam DeleteEcsImageLocal request param
type DeleteEcsImageLocalParam struct {
	BaseParam
	Params DeleteEcsImageLocalDetailParam `json:"params"`
}
