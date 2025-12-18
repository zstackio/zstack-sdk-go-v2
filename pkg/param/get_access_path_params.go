// Copyright (c) ZStack.io, Inc.

package param

// GetAccessPathDetailParam GetAccessPath detail param
type GetAccessPathDetailParam struct {
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
}

// GetAccessPathParam GetAccessPath request param
type GetAccessPathParam struct {
	BaseParam
	Params GetAccessPathDetailParam `json:"params"`
}
