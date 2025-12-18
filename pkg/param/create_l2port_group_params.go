// Copyright (c) ZStack.io, Inc.

package param

// CreateL2PortGroupDetailParam CreateL2PortGroup detail param
type CreateL2PortGroupDetailParam struct {
	VSwitchUuid string `json:"vSwitchUuid" validate:"required"`
	VlanMode string `json:"vlanMode,omitempty"`
	Vlan int `json:"vlan" validate:"required"`
	VlanRanges string `json:"vlanRanges,omitempty"`
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

// CreateL2PortGroupParam CreateL2PortGroup request param
type CreateL2PortGroupParam struct {
	BaseParam
	Params CreateL2PortGroupDetailParam `json:"params"`
}
