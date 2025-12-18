// Copyright (c) ZStack.io, Inc.

package param

// GetOssBucketFileFromRemoteDetailParam GetOssBucketFileFromRemote详细参数
type GetOssBucketFileFromRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"ossDomain,omitempty"`
	rest string `json:"ossKey,omitempty"`
	rest string `json:"ossSecret,omitempty"`
}

// GetOssBucketFileFromRemoteParam GetOssBucketFileFromRemote请求参数
type GetOssBucketFileFromRemoteParam struct {
	BaseParam
	Params GetOssBucketFileFromRemoteDetailParam `json:"params"` // 详细参数
}

