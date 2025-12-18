// Copyright (c) ZStack.io, Inc.

package param

// UpdateOssBucketDetailParam UpdateOssBucket详细参数
type UpdateOssBucketDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"ossDomain,omitempty"`
	rest string `json:"ossKey,omitempty"`
	rest string `json:"ossSecret,omitempty"`
}

// UpdateOssBucketParam UpdateOssBucket请求参数
type UpdateOssBucketParam struct {
	BaseParam
	Params UpdateOssBucketDetailParam `json:"params"` // 详细参数
}

