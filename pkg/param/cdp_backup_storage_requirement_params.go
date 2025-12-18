// Copyright (c) ZStack.io, Inc.

package param

// GetCdpBackupStorageRequirementDetailParam GetCdpBackupStorageRequirement详细参数
type GetCdpBackupStorageRequirementDetailParam struct {
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetCdpBackupStorageRequirementParam GetCdpBackupStorageRequirement请求参数
type GetCdpBackupStorageRequirementParam struct {
	BaseParam
	Params GetCdpBackupStorageRequirementDetailParam `json:"params"` // 详细参数
}

