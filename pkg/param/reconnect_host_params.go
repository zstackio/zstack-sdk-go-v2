// Copyright (c) ZStack.io, Inc.

package param

// ReconnectHostDetailParam ReconnectHost detail param
type ReconnectHostDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectHostParam ReconnectHost request param
type ReconnectHostParam struct {
	BaseParam
	Params ReconnectHostDetailParam `json:"params"`
}
