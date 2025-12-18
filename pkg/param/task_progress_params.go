// Copyright (c) ZStack.io, Inc.

package param

// GetTaskProgressDetailParam GetTaskProgress详细参数
type GetTaskProgressDetailParam struct {
	rest string `json:"apiId,omitempty"`
	rest bool `json:"all,omitempty"`
}

// GetTaskProgressParam GetTaskProgress请求参数
type GetTaskProgressParam struct {
	BaseParam
	Params GetTaskProgressDetailParam `json:"params"` // 详细参数
}

