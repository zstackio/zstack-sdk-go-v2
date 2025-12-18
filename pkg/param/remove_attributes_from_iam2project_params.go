// Copyright (c) ZStack.io, Inc.

package param

// RemoveAttributesFromIAM2ProjectDetailParam RemoveAttributesFromIAM2Project detail param
type RemoveAttributesFromIAM2ProjectDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	AttributeUuids []string `json:"attributeUuids" validate:"required"`
}

// RemoveAttributesFromIAM2ProjectParam RemoveAttributesFromIAM2Project request param
type RemoveAttributesFromIAM2ProjectParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2ProjectDetailParam `json:"params"`
}
