// Copyright (c) ZStack.io, Inc.

package param

// AddStorageProtocolDetailParam AddStorageProtocol detail param
type AddStorageProtocolDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	OutputProtocol string `json:"outputProtocol" validate:"required"`
}

// AddStorageProtocolParam AddStorageProtocol request param
type AddStorageProtocolParam struct {
	BaseParam
	Params AddStorageProtocolDetailParam `json:"params"`
}
