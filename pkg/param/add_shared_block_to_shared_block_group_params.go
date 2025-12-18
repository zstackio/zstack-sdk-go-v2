// Copyright (c) ZStack.io, Inc.

package param

// AddSharedBlockToSharedBlockGroupDetailParam AddSharedBlockToSharedBlockGroup detail param
type AddSharedBlockToSharedBlockGroupDetailParam struct {
	DiskUuid string `json:"diskUuid" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
}

// AddSharedBlockToSharedBlockGroupParam AddSharedBlockToSharedBlockGroup request param
type AddSharedBlockToSharedBlockGroupParam struct {
	BaseParam
	Params AddSharedBlockToSharedBlockGroupDetailParam `json:"params"`
}
