// Copyright (c) ZStack.io, Inc.

package param

// DeletePortMirrorSessionDetailParam DeletePortMirrorSession detail param
type DeletePortMirrorSessionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePortMirrorSessionParam DeletePortMirrorSession request param
type DeletePortMirrorSessionParam struct {
	BaseParam
	Params DeletePortMirrorSessionDetailParam `json:"params"`
}
