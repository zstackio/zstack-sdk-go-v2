// Copyright (c) ZStack.io, Inc.

package param

// SetImageQgaDetailParam SetImageQga详细参数
type SetImageQgaDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"enable" validate:"required"` // 必填
}

// SetImageQgaParam SetImageQga请求参数
type SetImageQgaParam struct {
	BaseParam
	Params SetImageQgaDetailParam `json:"params"` // 详细参数
}

