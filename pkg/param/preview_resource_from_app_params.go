// Copyright (c) ZStack.io, Inc.

package param

// PreviewResourceFromAppDetailParam PreviewResourceFromApp详细参数
type PreviewResourceFromAppDetailParam struct {
	rest string `json:"appUuid" validate:"required"` // 必填
	rest string `json:"parameters,omitempty"`
}

// PreviewResourceFromAppParam PreviewResourceFromApp请求参数
type PreviewResourceFromAppParam struct {
	BaseParam
	Params PreviewResourceFromAppDetailParam `json:"params"` // 详细参数
}

