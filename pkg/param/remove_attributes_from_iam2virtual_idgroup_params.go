// Copyright (c) ZStack.io, Inc.

package param

// RemoveAttributesFromIAM2VirtualIDGroupDetailParam RemoveAttributesFromIAM2VirtualIDGroup detail param
type RemoveAttributesFromIAM2VirtualIDGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	AttributeUuids []string `json:"attributeUuids" validate:"required"`
}

// RemoveAttributesFromIAM2VirtualIDGroupParam RemoveAttributesFromIAM2VirtualIDGroup request param
type RemoveAttributesFromIAM2VirtualIDGroupParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2VirtualIDGroupDetailParam `json:"params"`
}
