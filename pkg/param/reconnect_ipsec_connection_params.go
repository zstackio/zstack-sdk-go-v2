// Copyright (c) ZStack.io, Inc.

package param

// ReconnectIPsecConnectionDetailParam ReconnectIPsecConnection detail param
type ReconnectIPsecConnectionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectIPsecConnectionParam ReconnectIPsecConnection request param
type ReconnectIPsecConnectionParam struct {
	BaseParam
	Params ReconnectIPsecConnectionDetailParam `json:"params"`
}
