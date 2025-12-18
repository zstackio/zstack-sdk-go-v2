// Copyright (c) ZStack.io, Inc.

package param

// UpdateNasFileSystemDetailParam UpdateNasFileSystem详细参数
type UpdateNasFileSystemDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateNasFileSystemParam UpdateNasFileSystem请求参数
type UpdateNasFileSystemParam struct {
	BaseParam
	Params UpdateNasFileSystemDetailParam `json:"params"` // 详细参数
}

