// Copyright (c) ZStack.io, Inc.

package param

// GetOssBucketNameFromRemoteDetailParam GetOssBucketNameFromRemote详细参数
type GetOssBucketNameFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"ossDomain,omitempty"`
	rest string `json:"ossKey,omitempty"`
	rest string `json:"ossSecret,omitempty"`
}

// GetOssBucketNameFromRemoteParam GetOssBucketNameFromRemote请求参数
type GetOssBucketNameFromRemoteParam struct {
	BaseParam
	Params GetOssBucketNameFromRemoteDetailParam `json:"params"` // 详细参数
}

