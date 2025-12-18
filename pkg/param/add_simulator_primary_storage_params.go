// Copyright (c) ZStack.io, Inc.

package param

// AddSimulatorPrimaryStorageDetailParam AddSimulatorPrimaryStorage detail param
type AddSimulatorPrimaryStorageDetailParam struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSimulatorPrimaryStorageParam AddSimulatorPrimaryStorage request param
type AddSimulatorPrimaryStorageParam struct {
	BaseParam
	Params AddSimulatorPrimaryStorageDetailParam `json:"params"`
}
