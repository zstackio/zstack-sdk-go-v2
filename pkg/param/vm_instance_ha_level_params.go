// Copyright (c) ZStack.io, Inc.

package param

// GetVmInstanceHaLevelDetailParam GetVmInstanceHaLevel详细参数
type GetVmInstanceHaLevelDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmInstanceHaLevelParam GetVmInstanceHaLevel请求参数
type GetVmInstanceHaLevelParam struct {
	BaseParam
	Params GetVmInstanceHaLevelDetailParam `json:"params"` // 详细参数
}

