// Copyright (c) ZStack.io, Inc.

package param

// RerunLongJobDetailParam RerunLongJob detail param
type RerunLongJobDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RerunLongJobParam RerunLongJob request param
type RerunLongJobParam struct {
	BaseParam
	Params RerunLongJobDetailParam `json:"params"`
}
