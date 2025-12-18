// Copyright (c) ZStack.io, Inc.

package param

// UpdateCephPrimaryStorageMonDetailParam UpdateCephPrimaryStorageMon详细参数
type UpdateCephPrimaryStorageMonDetailParam struct {
	rest string `json:"monUuid" validate:"required"` // 必填
	rest string `json:"hostname,omitempty"`
	rest string `json:"sshUsername,omitempty"`
	rest string `json:"sshPassword,omitempty"`
	rest int `json:"sshPort,omitempty"`
	rest int `json:"monPort,omitempty"`
}

// UpdateCephPrimaryStorageMonParam UpdateCephPrimaryStorageMon请求参数
type UpdateCephPrimaryStorageMonParam struct {
	BaseParam
	Params UpdateCephPrimaryStorageMonDetailParam `json:"params"` // 详细参数
}

