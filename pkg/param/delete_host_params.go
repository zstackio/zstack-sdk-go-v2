// Copyright (c) ZStack.io, Inc.

package param

// DeleteHostDetailParam DeleteHost detail param
type DeleteHostDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHostParam DeleteHost request param
type DeleteHostParam struct {
	BaseParam
	Params DeleteHostDetailParam `json:"params"`
}
