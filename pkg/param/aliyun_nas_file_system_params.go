// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunNasFileSystemDetailParam CreateAliyunNasFileSystem详细参数
type CreateAliyunNasFileSystemDetailParam struct {
	rest string `json:"storageType" validate:"required"` // 必填
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"protocol,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateAliyunNasFileSystemParam CreateAliyunNasFileSystem请求参数
type CreateAliyunNasFileSystemParam struct {
	BaseParam
	Params CreateAliyunNasFileSystemDetailParam `json:"params"` // 详细参数
}

