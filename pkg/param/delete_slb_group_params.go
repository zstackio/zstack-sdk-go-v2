// Copyright (c) ZStack.io, Inc.

package param

// DeleteSlbGroupDetailParam DeleteSlbGroup detail param
type DeleteSlbGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSlbGroupParam DeleteSlbGroup request param
type DeleteSlbGroupParam struct {
	BaseParam
	Params DeleteSlbGroupDetailParam `json:"params"`
}
