// Copyright (c) ZStack.io, Inc.

package param

// SetVmQgaDetailParam SetVmQga详细参数
type SetVmQgaDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"enable" validate:"required"` // 必填
}

// SetVmQgaParam SetVmQga请求参数
type SetVmQgaParam struct {
	BaseParam
	Params SetVmQgaDetailParam `json:"params"` // 详细参数
}

