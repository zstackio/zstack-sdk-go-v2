// Copyright (c) ZStack.io, Inc.

package param

// AddAttributesToIAM2ProjectDetailParam AddAttributesToIAM2Project detail param
type AddAttributesToIAM2ProjectDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Attributes []interface{} `json:"attributes" validate:"required"`
}

// AddAttributesToIAM2ProjectParam AddAttributesToIAM2Project request param
type AddAttributesToIAM2ProjectParam struct {
	BaseParam
	Params AddAttributesToIAM2ProjectDetailParam `json:"params"`
}
