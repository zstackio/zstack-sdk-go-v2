// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelDetailParam DeleteModel detail param
type DeleteModelDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteModelParam DeleteModel request param
type DeleteModelParam struct {
	BaseParam
	Params DeleteModelDetailParam `json:"params"`
}
