// Copyright (c) ZStack.io, Inc.

package param

// DeleteResourceStackDetailParam DeleteResourceStack detail param
type DeleteResourceStackDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteResourceStackParam DeleteResourceStack request param
type DeleteResourceStackParam struct {
	BaseParam
	Params DeleteResourceStackDetailParam `json:"params"`
}
