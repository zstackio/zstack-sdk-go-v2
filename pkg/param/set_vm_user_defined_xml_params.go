// Copyright (c) ZStack.io, Inc.

package param

// SetVmUserDefinedXmlDetailParam SetVmUserDefinedXml detail param
type SetVmUserDefinedXmlDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	XmlBase64 string `json:"xmlBase64" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SetVmUserDefinedXmlParam SetVmUserDefinedXml request param
type SetVmUserDefinedXmlParam struct {
	BaseParam
	Params SetVmUserDefinedXmlDetailParam `json:"params"`
}
