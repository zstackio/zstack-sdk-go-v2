// Copyright (c) ZStack.io, Inc.

package param

// DisableCdpTaskDetailParam DisableCdpTask详细参数
type DisableCdpTaskDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"force,omitempty"`
}

// DisableCdpTaskParam DisableCdpTask请求参数
type DisableCdpTaskParam struct {
	BaseParam
	Params DisableCdpTaskDetailParam `json:"params"` // 详细参数
}

