// Copyright (c) ZStack.io, Inc.

package param

// ChangePortMirrorStateDetailParam ChangePortMirrorState detail param
type ChangePortMirrorStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangePortMirrorStateParam ChangePortMirrorState request param
type ChangePortMirrorStateParam struct {
	BaseParam
	Params ChangePortMirrorStateDetailParam `json:"params"`
}
