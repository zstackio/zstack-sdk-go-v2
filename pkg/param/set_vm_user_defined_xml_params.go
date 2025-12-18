// Copyright (c) ZStack.io, Inc.

package param

// SetVmUserDefinedXmlDetailParam SetVmUserDefinedXml详细参数
type SetVmUserDefinedXmlDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"xmlBase64" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SetVmUserDefinedXmlParam SetVmUserDefinedXml请求参数
type SetVmUserDefinedXmlParam struct {
	BaseParam
	Params SetVmUserDefinedXmlDetailParam `json:"params"` // 详细参数
}

