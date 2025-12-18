// Copyright (c) ZStack.io, Inc.

package param

// GetL3NetworkMtuDetailParam GetL3NetworkMtu detail param
type GetL3NetworkMtuDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
}

// GetL3NetworkMtuParam GetL3NetworkMtu request param
type GetL3NetworkMtuParam struct {
	BaseParam
	Params GetL3NetworkMtuDetailParam `json:"params"`
}
