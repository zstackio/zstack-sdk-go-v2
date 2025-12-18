// Copyright (c) ZStack.io, Inc.

package param

// ParseOvfDetailParam ParseOvf详细参数
type ParseOvfDetailParam struct {
	rest string `json:"xmlBase64" validate:"required"` // 必填
}

// ParseOvfParam ParseOvf请求参数
type ParseOvfParam struct {
	BaseParam
	Params ParseOvfDetailParam `json:"params"` // 详细参数
}

