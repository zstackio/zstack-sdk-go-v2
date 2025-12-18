// Copyright (c) ZStack.io, Inc.

package param

// GetVmXmlDetailParam GetVmXml detail param
type GetVmXmlDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// GetVmXmlParam GetVmXml request param
type GetVmXmlParam struct {
	BaseParam
	Params GetVmXmlDetailParam `json:"params"`
}
