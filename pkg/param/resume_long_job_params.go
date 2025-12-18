// Copyright (c) ZStack.io, Inc.

package param

// ResumeLongJobDetailParam ResumeLongJob详细参数
type ResumeLongJobDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ResumeLongJobParam ResumeLongJob请求参数
type ResumeLongJobParam struct {
	BaseParam
	Params ResumeLongJobDetailParam `json:"params"` // 详细参数
}

