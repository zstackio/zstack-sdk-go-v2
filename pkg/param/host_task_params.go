// Copyright (c) ZStack.io, Inc.

package param

// GetHostTaskDetailParam GetHostTask详细参数
type GetHostTaskDetailParam struct {
	rest []string `json:"hostUuids" validate:"required"` // 必填
	rest []string `json:"syncSignatures,omitempty"`
}

// GetHostTaskParam GetHostTask请求参数
type GetHostTaskParam struct {
	BaseParam
	Params GetHostTaskDetailParam `json:"params"` // 详细参数
}

