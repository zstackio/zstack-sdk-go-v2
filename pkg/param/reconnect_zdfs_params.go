// Copyright (c) ZStack.io, Inc.

package param

// ReconnectZdfsDetailParam ReconnectZdfs详细参数
type ReconnectZdfsDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectZdfsParam ReconnectZdfs请求参数
type ReconnectZdfsParam struct {
	BaseParam
	Params ReconnectZdfsDetailParam `json:"params"` // 详细参数
}

