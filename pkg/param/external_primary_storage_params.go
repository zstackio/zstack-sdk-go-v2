// Copyright (c) ZStack.io, Inc.

package param

// UpdateExternalPrimaryStorageDetailParam UpdateExternalPrimaryStorage详细参数
type UpdateExternalPrimaryStorageDetailParam struct {
	rest string `json:"config,omitempty"`
	rest string `json:"defaultProtocol,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"url,omitempty"`
}

// UpdateExternalPrimaryStorageParam UpdateExternalPrimaryStorage请求参数
type UpdateExternalPrimaryStorageParam struct {
	BaseParam
	Params UpdateExternalPrimaryStorageDetailParam `json:"params"` // 详细参数
}

