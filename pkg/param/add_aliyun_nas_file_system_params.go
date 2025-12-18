// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunNasFileSystemDetailParam AddAliyunNasFileSystem detail param
type AddAliyunNasFileSystemDetailParam struct {
	FileSystemId string `json:"fileSystemId" validate:"required"`
	Name string `json:"name" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasFileSystemParam AddAliyunNasFileSystem request param
type AddAliyunNasFileSystemParam struct {
	BaseParam
	Params AddAliyunNasFileSystemDetailParam `json:"params"`
}
