// Copyright (c) ZStack.io, Inc.

package param

// RemoveAttributesFromIAM2VirtualIDDetailParam RemoveAttributesFromIAM2VirtualID detail param
type RemoveAttributesFromIAM2VirtualIDDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	AttributeUuids []string `json:"attributeUuids" validate:"required"`
}

// RemoveAttributesFromIAM2VirtualIDParam RemoveAttributesFromIAM2VirtualID request param
type RemoveAttributesFromIAM2VirtualIDParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2VirtualIDDetailParam `json:"params"`
}
