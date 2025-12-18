// Copyright (c) ZStack.io, Inc.

package param

// ReconnectPrimaryStorageDetailParam ReconnectPrimaryStorage详细参数
type ReconnectPrimaryStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectPrimaryStorageParam ReconnectPrimaryStorage请求参数
type ReconnectPrimaryStorageParam struct {
	BaseParam
	Params ReconnectPrimaryStorageDetailParam `json:"params"` // 详细参数
}

