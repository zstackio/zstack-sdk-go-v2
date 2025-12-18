// Copyright (c) ZStack.io, Inc.

package param

// GetVmQgaDetailParam GetVmQga详细参数
type GetVmQgaDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmQgaParam GetVmQga请求参数
type GetVmQgaParam struct {
	BaseParam
	Params GetVmQgaDetailParam `json:"params"` // 详细参数
}

