// Copyright (c) ZStack.io, Inc.

package param

// ReclaimSpaceFromImageStoreDetailParam ReclaimSpaceFromImageStore详细参数
type ReclaimSpaceFromImageStoreDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReclaimSpaceFromImageStoreParam ReclaimSpaceFromImageStore请求参数
type ReclaimSpaceFromImageStoreParam struct {
	BaseParam
	Params ReclaimSpaceFromImageStoreDetailParam `json:"params"` // 详细参数
}

