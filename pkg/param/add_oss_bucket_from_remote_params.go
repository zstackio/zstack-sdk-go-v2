// Copyright (c) ZStack.io, Inc.

package param

// AddOssBucketFromRemoteDetailParam AddOssBucketFromRemote详细参数
type AddOssBucketFromRemoteDetailParam struct {
	rest string `json:"bucketName" validate:"required"` // 必填
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"ossDomain,omitempty"`
	rest string `json:"ossKey,omitempty"`
	rest string `json:"ossSecret,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddOssBucketFromRemoteParam AddOssBucketFromRemote请求参数
type AddOssBucketFromRemoteParam struct {
	BaseParam
	Params AddOssBucketFromRemoteDetailParam `json:"params"` // 详细参数
}

