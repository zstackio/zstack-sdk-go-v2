// Copyright (c) ZStack.io, Inc.

package param

// DeleteEipDetailParam DeleteEip detail param
type DeleteEipDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEipParam DeleteEip request param
type DeleteEipParam struct {
	BaseParam
	Params DeleteEipDetailParam `json:"params"`
}
