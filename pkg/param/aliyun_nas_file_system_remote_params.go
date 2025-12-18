// Copyright (c) ZStack.io, Inc.

package param

// GetAliyunNasFileSystemRemoteDetailParam GetAliyunNasFileSystemRemote详细参数
type GetAliyunNasFileSystemRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"fileSystemId,omitempty"`
}

// GetAliyunNasFileSystemRemoteParam GetAliyunNasFileSystemRemote请求参数
type GetAliyunNasFileSystemRemoteParam struct {
	BaseParam
	Params GetAliyunNasFileSystemRemoteDetailParam `json:"params"` // 详细参数
}

