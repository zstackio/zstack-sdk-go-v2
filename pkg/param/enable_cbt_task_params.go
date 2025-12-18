// Copyright (c) ZStack.io, Inc.

package param

// EnableCbtTaskDetailParam EnableCbtTask详细参数
type EnableCbtTaskDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"bitmapName,omitempty"`
}

// EnableCbtTaskParam EnableCbtTask请求参数
type EnableCbtTaskParam struct {
	BaseParam
	Params EnableCbtTaskDetailParam `json:"params"` // 详细参数
}

