// Copyright (c) ZStack.io, Inc.

package param

// DeleteOssBucketFileRemoteDetailParam DeleteOssBucketFileRemote详细参数
type DeleteOssBucketFileRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"fileName" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteOssBucketFileRemoteParam DeleteOssBucketFileRemote请求参数
type DeleteOssBucketFileRemoteParam struct {
	BaseParam
	Params DeleteOssBucketFileRemoteDetailParam `json:"params"` // 详细参数
}

