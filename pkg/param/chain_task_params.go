// Copyright (c) ZStack.io, Inc.

package param

// GetChainTaskDetailParam GetChainTask详细参数
type GetChainTaskDetailParam struct {
	rest []string `json:"syncSignatures,omitempty"`
}

// GetChainTaskParam GetChainTask请求参数
type GetChainTaskParam struct {
	BaseParam
	Params GetChainTaskDetailParam `json:"params"` // 详细参数
}

