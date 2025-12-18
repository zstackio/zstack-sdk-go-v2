// Copyright (c) ZStack.io, Inc.

package param

// AddStorageProtocolDetailParam AddStorageProtocol详细参数
type AddStorageProtocolDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"outputProtocol" validate:"required"` // 必填
}

// AddStorageProtocolParam AddStorageProtocol请求参数
type AddStorageProtocolParam struct {
	BaseParam
	Params AddStorageProtocolDetailParam `json:"params"` // 详细参数
}

