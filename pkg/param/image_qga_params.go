// Copyright (c) ZStack.io, Inc.

package param

// GetImageQgaDetailParam GetImageQga详细参数
type GetImageQgaDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetImageQgaParam GetImageQga请求参数
type GetImageQgaParam struct {
	BaseParam
	Params GetImageQgaDetailParam `json:"params"` // 详细参数
}

