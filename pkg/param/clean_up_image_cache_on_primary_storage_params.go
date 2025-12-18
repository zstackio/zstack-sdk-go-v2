// Copyright (c) ZStack.io, Inc.

package param

// CleanUpImageCacheOnPrimaryStorageDetailParam CleanUpImageCacheOnPrimaryStorage detail param
type CleanUpImageCacheOnPrimaryStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Force bool `json:"force,omitempty"`
}

// CleanUpImageCacheOnPrimaryStorageParam CleanUpImageCacheOnPrimaryStorage request param
type CleanUpImageCacheOnPrimaryStorageParam struct {
	BaseParam
	Params CleanUpImageCacheOnPrimaryStorageDetailParam `json:"params"`
}
