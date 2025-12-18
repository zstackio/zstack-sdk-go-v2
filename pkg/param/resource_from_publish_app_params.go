// Copyright (c) ZStack.io, Inc.

package param

// GetResourceFromPublishAppDetailParam GetResourceFromPublishApp详细参数
type GetResourceFromPublishAppDetailParam struct {
	rest string `json:"uuid,omitempty"`
}

// GetResourceFromPublishAppParam GetResourceFromPublishApp请求参数
type GetResourceFromPublishAppParam struct {
	BaseParam
	Params GetResourceFromPublishAppDetailParam `json:"params"` // 详细参数
}

