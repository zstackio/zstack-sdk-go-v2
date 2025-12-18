// Copyright (c) ZStack.io, Inc.

package param

// SetVmNumaDetailParam SetVmNuma详细参数
type SetVmNumaDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"enable" validate:"required"` // 必填
}

// SetVmNumaParam SetVmNuma请求参数
type SetVmNumaParam struct {
	BaseParam
	Params SetVmNumaDetailParam `json:"params"` // 详细参数
}

