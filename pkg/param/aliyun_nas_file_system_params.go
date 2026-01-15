// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddAliyunNasFileSystemParamDetail AddAliyunNasFileSystem detail param
type AddAliyunNasFileSystemParamDetail struct {
	FileSystemId string `json:"fileSystemId" validate:"required"`
	Name string `json:"name" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasFileSystemParam AddAliyunNasFileSystem request param
type AddAliyunNasFileSystemParam struct {
	BaseParam
	Params AddAliyunNasFileSystemParamDetail `json:"addAliyunNasFileSystem"`
}
// CreateAliyunNasFileSystemParamDetail CreateAliyunNasFileSystem detail param
type CreateAliyunNasFileSystemParamDetail struct {
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
	Params CreateAliyunNasFileSystemParamDetail `json:"createAliyunNasFileSystem"`
}
