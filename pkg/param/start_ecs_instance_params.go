// Copyright (c) ZStack.io, Inc.

package param

// StartEcsInstanceDetailParam StartEcsInstance detail param
type StartEcsInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StartEcsInstanceParam StartEcsInstance request param
type StartEcsInstanceParam struct {
	BaseParam
	Params StartEcsInstanceDetailParam `json:"params"`
}
