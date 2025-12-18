// Copyright (c) ZStack.io, Inc.

package param

// SetVmRDPDetailParam SetVmRDP详细参数
type SetVmRDPDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"enable" validate:"required"` // 必填
}

// SetVmRDPParam SetVmRDP请求参数
type SetVmRDPParam struct {
	BaseParam
	Params SetVmRDPDetailParam `json:"params"` // 详细参数
}

