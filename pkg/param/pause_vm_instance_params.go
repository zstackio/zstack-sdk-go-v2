// Copyright (c) ZStack.io, Inc.

package param

// PauseVmInstanceDetailParam PauseVmInstance详细参数
type PauseVmInstanceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// PauseVmInstanceParam PauseVmInstance请求参数
type PauseVmInstanceParam struct {
	BaseParam
	Params PauseVmInstanceDetailParam `json:"params"` // 详细参数
}

