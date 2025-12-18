// Copyright (c) ZStack.io, Inc.

package param

// ReclaimSpaceFromImageStoreDetailParam ReclaimSpaceFromImageStore detail param
type ReclaimSpaceFromImageStoreDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReclaimSpaceFromImageStoreParam ReclaimSpaceFromImageStore request param
type ReclaimSpaceFromImageStoreParam struct {
	BaseParam
	Params ReclaimSpaceFromImageStoreDetailParam `json:"params"`
}
