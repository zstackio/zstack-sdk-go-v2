// Copyright (c) ZStack.io, Inc.

package param

// GetOssBucketFileFromRemoteDetailParam GetOssBucketFileFromRemote detail param
type GetOssBucketFileFromRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	OssDomain string `json:"ossDomain,omitempty"`
	OssKey string `json:"ossKey,omitempty"`
	OssSecret string `json:"ossSecret,omitempty"`
}

// GetOssBucketFileFromRemoteParam GetOssBucketFileFromRemote request param
type GetOssBucketFileFromRemoteParam struct {
	BaseParam
	Params GetOssBucketFileFromRemoteDetailParam `json:"params"`
}
