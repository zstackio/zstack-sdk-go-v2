// Copyright (c) ZStack.io, Inc.

package param

// AddXDragonHostDetailParam AddXDragonHost详细参数
type AddXDragonHostDetailParam struct {
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest int `json:"cpuNum,omitempty"`
	rest int `json:"cpuSockets,omitempty"`
	rest int64 `json:"totalPhysicalMemory,omitempty"`
	rest int `json:"sshPort,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddXDragonHostParam AddXDragonHost请求参数
type AddXDragonHostParam struct {
	BaseParam
	Params AddXDragonHostDetailParam `json:"params"` // 详细参数
}

