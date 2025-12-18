// Copyright (c) ZStack.io, Inc.

package param

// GetVmRDPDetailParam GetVmRDP详细参数
type GetVmRDPDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmRDPParam GetVmRDP请求参数
type GetVmRDPParam struct {
	BaseParam
	Params GetVmRDPDetailParam `json:"params"` // 详细参数
}

