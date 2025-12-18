// Copyright (c) ZStack.io, Inc.

package param

// CreateOssBucketRemoteDetailParam CreateOssBucketRemote详细参数
type CreateOssBucketRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"bucketName" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"ossDomain,omitempty"`
	rest string `json:"ossKey,omitempty"`
	rest string `json:"ossSecret,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateOssBucketRemoteParam CreateOssBucketRemote请求参数
type CreateOssBucketRemoteParam struct {
	BaseParam
	Params CreateOssBucketRemoteDetailParam `json:"params"` // 详细参数
}

