// Copyright (c) ZStack.io, Inc.

package param

// BatchQueryDetailParam BatchQuery详细参数
type BatchQueryDetailParam struct {
	rest string `json:"script,omitempty"`
}

// BatchQueryParam BatchQuery请求参数
type BatchQueryParam struct {
	BaseParam
	Params BatchQueryDetailParam `json:"params"` // 详细参数
}

