// Copyright (c) ZStack.io, Inc.

package param

// SetL3NetworkMtuDetailParam SetL3NetworkMtu detail param
type SetL3NetworkMtuDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Mtu int `json:"mtu" validate:"required"`
}

// SetL3NetworkMtuParam SetL3NetworkMtu request param
type SetL3NetworkMtuParam struct {
	BaseParam
	Params SetL3NetworkMtuDetailParam `json:"params"`
}
