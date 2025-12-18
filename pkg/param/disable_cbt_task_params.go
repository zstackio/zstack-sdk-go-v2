// Copyright (c) ZStack.io, Inc.

package param

// DisableCbtTaskDetailParam DisableCbtTask详细参数
type DisableCbtTaskDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"force,omitempty"`
}

// DisableCbtTaskParam DisableCbtTask请求参数
type DisableCbtTaskParam struct {
	BaseParam
	Params DisableCbtTaskDetailParam `json:"params"` // 详细参数
}

