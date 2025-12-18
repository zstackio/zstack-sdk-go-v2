// Copyright (c) ZStack.io, Inc.

package param

// SetImageBootModeDetailParam SetImageBootMode detail param
type SetImageBootModeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BootMode string `json:"bootMode" validate:"required"`
}

// SetImageBootModeParam SetImageBootMode request param
type SetImageBootModeParam struct {
	BaseParam
	Params SetImageBootModeDetailParam `json:"params"`
}
