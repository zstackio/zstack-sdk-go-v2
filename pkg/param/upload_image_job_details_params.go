// Copyright (c) ZStack.io, Inc.

package param

// GetUploadImageJobDetailsDetailParam GetUploadImageJobDetails详细参数
type GetUploadImageJobDetailsDetailParam struct {
	rest string `json:"imageId" validate:"required"` // 必填
}

// GetUploadImageJobDetailsParam GetUploadImageJobDetails请求参数
type GetUploadImageJobDetailsParam struct {
	BaseParam
	Params GetUploadImageJobDetailsDetailParam `json:"params"` // 详细参数
}

