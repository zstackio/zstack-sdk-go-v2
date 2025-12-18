// Copyright (c) ZStack.io, Inc.

package param

// CreateL2HardwareVxlanNetworkDetailParam CreateL2HardwareVxlanNetwork detail param
type CreateL2HardwareVxlanNetworkDetailParam struct {
	Vni int `json:"vni,omitempty"`
	PoolUuid string `json:"poolUuid" validate:"required"`
	H3cTenantUuid string `json:"h3cTenantUuid,omitempty"`
	Vlan int `json:"vlan,omitempty"`
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

// CreateL2HardwareVxlanNetworkParam CreateL2HardwareVxlanNetwork request param
type CreateL2HardwareVxlanNetworkParam struct {
	BaseParam
	Params CreateL2HardwareVxlanNetworkDetailParam `json:"params"`
}
