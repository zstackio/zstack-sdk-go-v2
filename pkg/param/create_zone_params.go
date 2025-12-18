// Copyright (c) ZStack.io, Inc.

package param

// CreateZoneDetailParam CreateZone detail param
type CreateZoneDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateZoneParam CreateZone request param
type CreateZoneParam struct {
	BaseParam
	Params CreateZoneDetailParam `json:"params"`
}
