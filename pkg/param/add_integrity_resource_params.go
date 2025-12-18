// Copyright (c) ZStack.io, Inc.

package param

// AddIntegrityResourceDetailParam AddIntegrityResource detail param
type AddIntegrityResourceDetailParam struct {
	ResourceType string `json:"resourceType" validate:"required"`
	IntegrityResourceDataRangeInDays int `json:"integrityResourceDataRangeInDays,omitempty"`
}

// AddIntegrityResourceParam AddIntegrityResource request param
type AddIntegrityResourceParam struct {
	BaseParam
	Params AddIntegrityResourceDetailParam `json:"params"`
}
