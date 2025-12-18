// Copyright (c) ZStack.io, Inc.

package param

// AddAttributesToIAM2VirtualIDGroupDetailParam AddAttributesToIAM2VirtualIDGroup detail param
type AddAttributesToIAM2VirtualIDGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Attributes []interface{} `json:"attributes" validate:"required"`
}

// AddAttributesToIAM2VirtualIDGroupParam AddAttributesToIAM2VirtualIDGroup request param
type AddAttributesToIAM2VirtualIDGroupParam struct {
	BaseParam
	Params AddAttributesToIAM2VirtualIDGroupDetailParam `json:"params"`
}
