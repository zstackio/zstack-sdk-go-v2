// Copyright (c) ZStack.io, Inc.

package param

// DeleteCbtTaskDetailParam DeleteCbtTask详细参数
type DeleteCbtTaskDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"force,omitempty"`
	rest string `json:"deleteMode,omitempty"`
}

// DeleteCbtTaskParam DeleteCbtTask请求参数
type DeleteCbtTaskParam struct {
	BaseParam
	Params DeleteCbtTaskDetailParam `json:"params"` // 详细参数
}

