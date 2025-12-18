// Copyright (c) ZStack.io, Inc.

package param

// ExpungeIAM2ProjectDetailParam ExpungeIAM2Project detail param
type ExpungeIAM2ProjectDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExpungeIAM2ProjectParam ExpungeIAM2Project request param
type ExpungeIAM2ProjectParam struct {
	BaseParam
	Params ExpungeIAM2ProjectDetailParam `json:"params"`
}
