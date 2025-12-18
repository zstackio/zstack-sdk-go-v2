// Copyright (c) ZStack.io, Inc.

package param

// GetVmTaskDetailParam GetVmTask详细参数
type GetVmTaskDetailParam struct {
	rest []string `json:"vmInstanceUuids" validate:"required"` // 必填
	rest []string `json:"syncSignatures,omitempty"`
}

// GetVmTaskParam GetVmTask请求参数
type GetVmTaskParam struct {
	BaseParam
	Params GetVmTaskDetailParam `json:"params"` // 详细参数
}

