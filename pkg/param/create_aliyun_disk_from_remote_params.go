// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunDiskFromRemoteDetailParam CreateAliyunDiskFromRemote detail param
type CreateAliyunDiskFromRemoteDetailParam struct {
	IdentityUuid string `json:"identityUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	SizeWithGB int `json:"sizeWithGB,omitempty"`
	Description string `json:"description,omitempty"`
	DiskCategory string `json:"diskCategory,omitempty"`
	SnapshotUuid string `json:"snapshotUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunDiskFromRemoteParam CreateAliyunDiskFromRemote request param
type CreateAliyunDiskFromRemoteParam struct {
	BaseParam
	Params CreateAliyunDiskFromRemoteDetailParam `json:"params"`
}
