// Copyright (c) ZStack.io, Inc.

package param

// UpdateCephBackupStorageMonDetailParam UpdateCephBackupStorageMon详细参数
type UpdateCephBackupStorageMonDetailParam struct {
	rest string `json:"monUuid" validate:"required"` // 必填
	rest string `json:"hostname,omitempty"`
	rest string `json:"sshUsername,omitempty"`
	rest string `json:"sshPassword,omitempty"`
	rest int `json:"sshPort,omitempty"`
	rest int `json:"monPort,omitempty"`
}

// UpdateCephBackupStorageMonParam UpdateCephBackupStorageMon请求参数
type UpdateCephBackupStorageMonParam struct {
	BaseParam
	Params UpdateCephBackupStorageMonDetailParam `json:"params"` // 详细参数
}

