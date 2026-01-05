// Copyright (c) ZStack.io, Inc.

package param

// DeleteNfvInstGroupDetailParam DeleteNfvInstGroup detail param
type DeleteNfvInstGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteNfvInstGroupParam DeleteNfvInstGroup request param
type DeleteNfvInstGroupParam struct {
	BaseParam
	Params DeleteNfvInstGroupDetailParam `json:"params"`
}
