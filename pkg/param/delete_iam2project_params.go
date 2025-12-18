// Copyright (c) ZStack.io, Inc.

package param

// DeleteIAM2ProjectDetailParam DeleteIAM2Project detail param
type DeleteIAM2ProjectDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteIAM2ProjectParam DeleteIAM2Project request param
type DeleteIAM2ProjectParam struct {
	BaseParam
	Params DeleteIAM2ProjectDetailParam `json:"params"`
}
