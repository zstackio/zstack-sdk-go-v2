// Copyright (c) ZStack.io, Inc.

package param

// UpdateZoneDetailParam UpdateZone detail param
type UpdateZoneDetailParam struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	IsDefault bool `json:"isDefault,omitempty"`
}

// UpdateZoneParam UpdateZone request param
type UpdateZoneParam struct {
	BaseParam
	Params UpdateZoneDetailParam `json:"params"`
}
