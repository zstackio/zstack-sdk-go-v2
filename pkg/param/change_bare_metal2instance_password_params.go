// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2InstancePasswordDetailParam ChangeBareMetal2InstancePassword detail param
type ChangeBareMetal2InstancePasswordDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ChangeBareMetal2InstancePasswordParam ChangeBareMetal2InstancePassword request param
type ChangeBareMetal2InstancePasswordParam struct {
	BaseParam
	Params ChangeBareMetal2InstancePasswordDetailParam `json:"params"`
}
