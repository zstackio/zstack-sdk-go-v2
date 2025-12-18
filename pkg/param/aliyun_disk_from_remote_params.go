// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunDiskFromRemoteDetailParam CreateAliyunDiskFromRemote详细参数
type CreateAliyunDiskFromRemoteDetailParam struct {
	rest string `json:"identityUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest int `json:"sizeWithGB,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"diskCategory,omitempty"`
	rest string `json:"snapshotUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateAliyunDiskFromRemoteParam CreateAliyunDiskFromRemote请求参数
type CreateAliyunDiskFromRemoteParam struct {
	BaseParam
	Params CreateAliyunDiskFromRemoteDetailParam `json:"params"` // 详细参数
}

