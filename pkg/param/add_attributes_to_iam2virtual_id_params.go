// Copyright (c) ZStack.io, Inc.

package param

// AddAttributesToIAM2VirtualIDDetailParam AddAttributesToIAM2VirtualID detail param
type AddAttributesToIAM2VirtualIDDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Attributes []interface{} `json:"attributes" validate:"required"`
}

// AddAttributesToIAM2VirtualIDParam AddAttributesToIAM2VirtualID request param
type AddAttributesToIAM2VirtualIDParam struct {
	BaseParam
	Params AddAttributesToIAM2VirtualIDDetailParam `json:"params"`
}
