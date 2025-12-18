// Copyright (c) ZStack.io, Inc.

package param

// AddSimulatorHostDetailParam AddSimulatorHost详细参数
type AddSimulatorHostDetailParam struct {
	rest int64 `json:"memoryCapacity" validate:"required"` // 必填
	rest int64 `json:"cpuCapacity" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSimulatorHostParam AddSimulatorHost请求参数
type AddSimulatorHostParam struct {
	BaseParam
	Params AddSimulatorHostDetailParam `json:"params"` // 详细参数
}

