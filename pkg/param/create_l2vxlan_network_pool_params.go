// Copyright (c) ZStack.io, Inc.

package param

// CreateL2VxlanNetworkPoolDetailParam CreateL2VxlanNetworkPool detail param
type CreateL2VxlanNetworkPoolDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	Isolated bool `json:"isolated,omitempty"`
	Pvlan string `json:"pvlan,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2VxlanNetworkPoolParam CreateL2VxlanNetworkPool request param
type CreateL2VxlanNetworkPoolParam struct {
	BaseParam
	Params CreateL2VxlanNetworkPoolDetailParam `json:"params"`
}
