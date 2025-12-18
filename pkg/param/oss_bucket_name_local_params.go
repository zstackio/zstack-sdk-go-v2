// Copyright (c) ZStack.io, Inc.

package param

// DeleteOssBucketNameLocalDetailParam DeleteOssBucketNameLocal详细参数
type DeleteOssBucketNameLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteOssBucketNameLocalParam DeleteOssBucketNameLocal请求参数
type DeleteOssBucketNameLocalParam struct {
	BaseParam
	Params DeleteOssBucketNameLocalDetailParam `json:"params"` // 详细参数
}

