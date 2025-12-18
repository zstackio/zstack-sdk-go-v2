// Copyright (c) ZStack.io, Inc.

package param

// SyncImageSizeDetailParam SyncImageSize详细参数
type SyncImageSizeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// SyncImageSizeParam SyncImageSize请求参数
type SyncImageSizeParam struct {
	BaseParam
	Params SyncImageSizeDetailParam `json:"params"` // 详细参数
}

