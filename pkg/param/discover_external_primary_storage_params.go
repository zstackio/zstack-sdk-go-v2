// Copyright (c) ZStack.io, Inc.

package param

// DiscoverExternalPrimaryStorageDetailParam DiscoverExternalPrimaryStorage详细参数
type DiscoverExternalPrimaryStorageDetailParam struct {
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"identity,omitempty"`
	rest string `json:"config,omitempty"`
}

// DiscoverExternalPrimaryStorageParam DiscoverExternalPrimaryStorage请求参数
type DiscoverExternalPrimaryStorageParam struct {
	BaseParam
	Params DiscoverExternalPrimaryStorageDetailParam `json:"params"` // 详细参数
}

