// Copyright (c) ZStack.io, Inc.

package param

// RecoverIAM2ProjectDetailParam RecoverIAM2Project detail param
type RecoverIAM2ProjectDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RecoverIAM2ProjectParam RecoverIAM2Project request param
type RecoverIAM2ProjectParam struct {
	BaseParam
	Params RecoverIAM2ProjectDetailParam `json:"params"`
}
