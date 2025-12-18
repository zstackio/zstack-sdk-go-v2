// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunSnapshotRemoteDetailParam CreateAliyunSnapshotRemote详细参数
type CreateAliyunSnapshotRemoteDetailParam struct {
	rest string `json:"diskUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateAliyunSnapshotRemoteParam CreateAliyunSnapshotRemote请求参数
type CreateAliyunSnapshotRemoteParam struct {
	BaseParam
	Params CreateAliyunSnapshotRemoteDetailParam `json:"params"` // 详细参数
}

