// Copyright (c) ZStack.io, Inc.

package param

// GetResourceFromResourceStackDetailParam GetResourceFromResourceStack detail param
type GetResourceFromResourceStackDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetResourceFromResourceStackParam GetResourceFromResourceStack request param
type GetResourceFromResourceStackParam struct {
	BaseParam
	Params GetResourceFromResourceStackDetailParam `json:"params"`
}
