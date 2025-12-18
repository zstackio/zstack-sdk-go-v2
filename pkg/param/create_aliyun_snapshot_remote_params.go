// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunSnapshotRemoteDetailParam CreateAliyunSnapshotRemote detail param
type CreateAliyunSnapshotRemoteDetailParam struct {
	DiskUuid string `json:"diskUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunSnapshotRemoteParam CreateAliyunSnapshotRemote request param
type CreateAliyunSnapshotRemoteParam struct {
	BaseParam
	Params CreateAliyunSnapshotRemoteDetailParam `json:"params"`
}
