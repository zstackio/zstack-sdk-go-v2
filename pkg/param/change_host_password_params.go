// Copyright (c) ZStack.io, Inc.

package param

// ChangeHostPasswordDetailParam ChangeHostPassword detail param
type ChangeHostPasswordDetailParam struct {
	HostUuid string `json:"hostUuid" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ChangeHostPasswordParam ChangeHostPassword request param
type ChangeHostPasswordParam struct {
	BaseParam
	Params ChangeHostPasswordDetailParam `json:"params"`
}
