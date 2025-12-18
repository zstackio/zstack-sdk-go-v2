// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmUserDefinedXmlDetailParam DeleteVmUserDefinedXml详细参数
type DeleteVmUserDefinedXmlDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVmUserDefinedXmlParam DeleteVmUserDefinedXml请求参数
type DeleteVmUserDefinedXmlParam struct {
	BaseParam
	Params DeleteVmUserDefinedXmlDetailParam `json:"params"` // 详细参数
}

