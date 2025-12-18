// Copyright (c) ZStack.io, Inc.

package param

// FlattenVmInstanceDetailParam FlattenVmInstance详细参数
type FlattenVmInstanceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"full,omitempty"`
	rest bool `json:"dryRun,omitempty"`
}

// FlattenVmInstanceParam FlattenVmInstance请求参数
type FlattenVmInstanceParam struct {
	BaseParam
	Params FlattenVmInstanceDetailParam `json:"params"` // 详细参数
}

