// Copyright (c) ZStack.io, Inc.

package param

// StopEcsInstanceDetailParam StopEcsInstance detail param
type StopEcsInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StopEcsInstanceParam StopEcsInstance request param
type StopEcsInstanceParam struct {
	BaseParam
	Params StopEcsInstanceDetailParam `json:"params"`
}
