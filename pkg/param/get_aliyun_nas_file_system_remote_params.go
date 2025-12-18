// Copyright (c) ZStack.io, Inc.

package param

// GetAliyunNasFileSystemRemoteDetailParam GetAliyunNasFileSystemRemote detail param
type GetAliyunNasFileSystemRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	FileSystemId string `json:"fileSystemId,omitempty"`
}

// GetAliyunNasFileSystemRemoteParam GetAliyunNasFileSystemRemote request param
type GetAliyunNasFileSystemRemoteParam struct {
	BaseParam
	Params GetAliyunNasFileSystemRemoteDetailParam `json:"params"`
}
