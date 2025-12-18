// Copyright (c) ZStack.io, Inc.

package param

// UpdatePortMirrorDetailParam UpdatePortMirror detail param
type UpdatePortMirrorDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdatePortMirrorParam UpdatePortMirror request param
type UpdatePortMirrorParam struct {
	BaseParam
	Params UpdatePortMirrorDetailParam `json:"params"`
}
