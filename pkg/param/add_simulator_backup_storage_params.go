// Copyright (c) ZStack.io, Inc.

package param

// AddSimulatorBackupStorageDetailParam AddSimulatorBackupStorage detail param
type AddSimulatorBackupStorageDetailParam struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ImportImages bool `json:"importImages,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSimulatorBackupStorageParam AddSimulatorBackupStorage request param
type AddSimulatorBackupStorageParam struct {
	BaseParam
	Params AddSimulatorBackupStorageDetailParam `json:"params"`
}
