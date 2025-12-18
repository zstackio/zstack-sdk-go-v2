// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmUserDefinedXmlDetailParam DeleteVmUserDefinedXml detail param
type DeleteVmUserDefinedXmlDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVmUserDefinedXmlParam DeleteVmUserDefinedXml request param
type DeleteVmUserDefinedXmlParam struct {
	BaseParam
	Params DeleteVmUserDefinedXmlDetailParam `json:"params"`
}
