// Copyright (c) ZStack.io, Inc.

package param

// AddSimulatorBackupStorageDetailParam AddSimulatorBackupStorage详细参数
type AddSimulatorBackupStorageDetailParam struct {
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest bool `json:"importImages,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSimulatorBackupStorageParam AddSimulatorBackupStorage请求参数
type AddSimulatorBackupStorageParam struct {
	BaseParam
	Params AddSimulatorBackupStorageDetailParam `json:"params"` // 详细参数
}

