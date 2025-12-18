// Copyright (c) ZStack.io, Inc.

package param

// SetVmInstanceHaLevelDetailParam SetVmInstanceHaLevel详细参数
type SetVmInstanceHaLevelDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"level" validate:"required"` // 必填
}

// SetVmInstanceHaLevelParam SetVmInstanceHaLevel请求参数
type SetVmInstanceHaLevelParam struct {
	BaseParam
	Params SetVmInstanceHaLevelDetailParam `json:"params"` // 详细参数
}

