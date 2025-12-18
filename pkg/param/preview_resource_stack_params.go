// Copyright (c) ZStack.io, Inc.

package param

// PreviewResourceStackDetailParam PreviewResourceStack详细参数
type PreviewResourceStackDetailParam struct {
	rest string `json:"type,omitempty"`
	rest string `json:"templateContent,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"parameters,omitempty"`
	rest string `json:"preParameters,omitempty"`
}

// PreviewResourceStackParam PreviewResourceStack请求参数
type PreviewResourceStackParam struct {
	BaseParam
	Params PreviewResourceStackDetailParam `json:"params"` // 详细参数
}

