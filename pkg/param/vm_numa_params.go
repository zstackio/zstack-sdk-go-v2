// Copyright (c) ZStack.io, Inc.

package param

// GetVmNumaDetailParam GetVmNuma详细参数
type GetVmNumaDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmNumaParam GetVmNuma请求参数
type GetVmNumaParam struct {
	BaseParam
	Params GetVmNumaDetailParam `json:"params"` // 详细参数
}

