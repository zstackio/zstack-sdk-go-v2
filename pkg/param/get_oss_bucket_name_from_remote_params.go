// Copyright (c) ZStack.io, Inc.

package param

// GetOssBucketNameFromRemoteDetailParam GetOssBucketNameFromRemote detail param
type GetOssBucketNameFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	OssDomain string `json:"ossDomain,omitempty"`
	OssKey string `json:"ossKey,omitempty"`
	OssSecret string `json:"ossSecret,omitempty"`
}

// GetOssBucketNameFromRemoteParam GetOssBucketNameFromRemote request param
type GetOssBucketNameFromRemoteParam struct {
	BaseParam
	Params GetOssBucketNameFromRemoteDetailParam `json:"params"`
}
