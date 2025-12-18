// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunNasFileSystemDetailParam AddAliyunNasFileSystem详细参数
type AddAliyunNasFileSystemDetailParam struct {
	rest string `json:"fileSystemId" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasFileSystemParam AddAliyunNasFileSystem请求参数
type AddAliyunNasFileSystemParam struct {
	BaseParam
	Params AddAliyunNasFileSystemDetailParam `json:"params"` // 详细参数
}

