// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsInstanceVncPasswordDetailParam UpdateEcsInstanceVncPassword detail param
type UpdateEcsInstanceVncPasswordDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UpdateEcsInstanceVncPasswordParam UpdateEcsInstanceVncPassword request param
type UpdateEcsInstanceVncPasswordParam struct {
	BaseParam
	Params UpdateEcsInstanceVncPasswordDetailParam `json:"params"`
}
