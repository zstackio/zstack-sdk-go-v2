// Copyright (c) ZStack.io, Inc.

package param

// CreateAffinityGroupDetailParam CreateAffinityGroup detail param
type CreateAffinityGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Policy string `json:"policy,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	SubType string `json:"subType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAffinityGroupParam CreateAffinityGroup request param
type CreateAffinityGroupParam struct {
	BaseParam
	Params CreateAffinityGroupDetailParam `json:"params"`
}
