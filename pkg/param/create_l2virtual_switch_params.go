// Copyright (c) ZStack.io, Inc.

package param

// CreateL2VirtualSwitchDetailParam CreateL2VirtualSwitch detail param
type CreateL2VirtualSwitchDetailParam struct {
	IsDistributed bool `json:"isDistributed,omitempty"`
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

// CreateL2VirtualSwitchParam CreateL2VirtualSwitch request param
type CreateL2VirtualSwitchParam struct {
	BaseParam
	Params CreateL2VirtualSwitchDetailParam `json:"params"`
}
