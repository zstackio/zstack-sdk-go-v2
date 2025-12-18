// Copyright (c) ZStack.io, Inc.

package param

// AddSimulatorPrimaryStorageDetailParam AddSimulatorPrimaryStorage详细参数
type AddSimulatorPrimaryStorageDetailParam struct {
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest int64 `json:"availablePhysicalCapacity,omitempty"`
	rest int64 `json:"totalPhysicalCapacity,omitempty"`
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSimulatorPrimaryStorageParam AddSimulatorPrimaryStorage请求参数
type AddSimulatorPrimaryStorageParam struct {
	BaseParam
	Params AddSimulatorPrimaryStorageDetailParam `json:"params"` // 详细参数
}

