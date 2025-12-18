// Copyright (c) ZStack.io, Inc.

package param

// CreateL2HardwareVxlanNetworkPoolDetailParam CreateL2HardwareVxlanNetworkPool detail param
type CreateL2HardwareVxlanNetworkPoolDetailParam struct {
	SdnControllerUuid string `json:"sdnControllerUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	PhysicalInterface string `json:"physicalInterface" validate:"required"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	Isolated bool `json:"isolated,omitempty"`
	Pvlan string `json:"pvlan,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2HardwareVxlanNetworkPoolParam CreateL2HardwareVxlanNetworkPool request param
type CreateL2HardwareVxlanNetworkPoolParam struct {
	BaseParam
	Params CreateL2HardwareVxlanNetworkPoolDetailParam `json:"params"`
}
