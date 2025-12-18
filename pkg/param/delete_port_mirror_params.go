// Copyright (c) ZStack.io, Inc.

package param

// DeletePortMirrorDetailParam DeletePortMirror detail param
type DeletePortMirrorDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePortMirrorParam DeletePortMirror request param
type DeletePortMirrorParam struct {
	BaseParam
	Params DeletePortMirrorDetailParam `json:"params"`
}
