// Copyright (c) ZStack.io, Inc.

package param

// GetEcsInstanceVncUrlDetailParam GetEcsInstanceVncUrl detail param
type GetEcsInstanceVncUrlDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetEcsInstanceVncUrlParam GetEcsInstanceVncUrl request param
type GetEcsInstanceVncUrlParam struct {
	BaseParam
	Params GetEcsInstanceVncUrlDetailParam `json:"params"`
}
