// Copyright (c) ZStack.io, Inc.

package param

// DeleteSlbGroupDetailParam DeleteSlbGroup详细参数
type DeleteSlbGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteSlbGroupParam DeleteSlbGroup请求参数
type DeleteSlbGroupParam struct {
	BaseParam
	Params DeleteSlbGroupDetailParam `json:"params"` // 详细参数
}

