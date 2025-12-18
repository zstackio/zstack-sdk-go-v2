// Copyright (c) ZStack.io, Inc.

package param

// DeleteMediaDetailParam DeleteMedia detail param
type DeleteMediaDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMediaParam DeleteMedia request param
type DeleteMediaParam struct {
	BaseParam
	Params DeleteMediaDetailParam `json:"params"`
}
