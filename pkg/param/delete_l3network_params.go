// Copyright (c) ZStack.io, Inc.

package param

// DeleteL3NetworkDetailParam DeleteL3Network detail param
type DeleteL3NetworkDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteL3NetworkParam DeleteL3Network request param
type DeleteL3NetworkParam struct {
	BaseParam
	Params DeleteL3NetworkDetailParam `json:"params"`
}
