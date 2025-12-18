// Copyright (c) ZStack.io, Inc.

package param

// GetVmXmlDetailParam GetVmXml详细参数
type GetVmXmlDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// GetVmXmlParam GetVmXml请求参数
type GetVmXmlParam struct {
	BaseParam
	Params GetVmXmlDetailParam `json:"params"` // 详细参数
}

