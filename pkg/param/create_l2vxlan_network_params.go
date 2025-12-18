// Copyright (c) ZStack.io, Inc.

package param

// CreateL2VxlanNetworkDetailParam CreateL2VxlanNetwork detail param
type CreateL2VxlanNetworkDetailParam struct {
	Vni int `json:"vni,omitempty"`
	PoolUuid string `json:"poolUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	Isolated bool `json:"isolated,omitempty"`
	Pvlan string `json:"pvlan,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2VxlanNetworkParam CreateL2VxlanNetwork request param
type CreateL2VxlanNetworkParam struct {
	BaseParam
	Params CreateL2VxlanNetworkDetailParam `json:"params"`
}
