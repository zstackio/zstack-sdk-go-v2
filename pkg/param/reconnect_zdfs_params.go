// Copyright (c) ZStack.io, Inc.

package param

// ReconnectZdfsDetailParam ReconnectZdfs detail param
type ReconnectZdfsDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectZdfsParam ReconnectZdfs request param
type ReconnectZdfsParam struct {
	BaseParam
	Params ReconnectZdfsDetailParam `json:"params"`
}
