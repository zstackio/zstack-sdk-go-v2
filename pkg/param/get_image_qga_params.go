// Copyright (c) ZStack.io, Inc.

package param

// GetImageQgaDetailParam GetImageQga detail param
type GetImageQgaDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetImageQgaParam GetImageQga request param
type GetImageQgaParam struct {
	BaseParam
	Params GetImageQgaDetailParam `json:"params"`
}
