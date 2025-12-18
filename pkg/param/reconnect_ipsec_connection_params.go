// Copyright (c) ZStack.io, Inc.

package param

// ReconnectIPsecConnectionDetailParam ReconnectIPsecConnection详细参数
type ReconnectIPsecConnectionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectIPsecConnectionParam ReconnectIPsecConnection请求参数
type ReconnectIPsecConnectionParam struct {
	BaseParam
	Params ReconnectIPsecConnectionDetailParam `json:"params"` // 详细参数
}

