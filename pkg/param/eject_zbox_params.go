// Copyright (c) ZStack.io, Inc.

package param

// EjectZBoxDetailParam EjectZBox detail param
type EjectZBoxDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// EjectZBoxParam EjectZBox request param
type EjectZBoxParam struct {
	BaseParam
	Params EjectZBoxDetailParam `json:"params"`
}
