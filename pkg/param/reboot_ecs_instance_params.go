// Copyright (c) ZStack.io, Inc.

package param

// RebootEcsInstanceDetailParam RebootEcsInstance detail param
type RebootEcsInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RebootEcsInstanceParam RebootEcsInstance request param
type RebootEcsInstanceParam struct {
	BaseParam
	Params RebootEcsInstanceDetailParam `json:"params"`
}
