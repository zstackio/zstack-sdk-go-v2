// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunNasFileSystemDetailParam CreateAliyunNasFileSystem detail param
type CreateAliyunNasFileSystemDetailParam struct {
	StorageType string `json:"storageType" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	Protocol string `json:"protocol,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunNasFileSystemParam CreateAliyunNasFileSystem request param
type CreateAliyunNasFileSystemParam struct {
	BaseParam
	Params CreateAliyunNasFileSystemDetailParam `json:"params"`
}
