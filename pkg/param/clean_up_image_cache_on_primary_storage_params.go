// Copyright (c) ZStack.io, Inc.

package param

// CleanUpImageCacheOnPrimaryStorageDetailParam CleanUpImageCacheOnPrimaryStorage详细参数
type CleanUpImageCacheOnPrimaryStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"force,omitempty"`
}

// CleanUpImageCacheOnPrimaryStorageParam CleanUpImageCacheOnPrimaryStorage请求参数
type CleanUpImageCacheOnPrimaryStorageParam struct {
	BaseParam
	Params CleanUpImageCacheOnPrimaryStorageDetailParam `json:"params"` // 详细参数
}

