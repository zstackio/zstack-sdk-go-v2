// Copyright (c) ZStack.io, Inc.

package param

// PrimaryStorageMigrateVmDetailParam PrimaryStorageMigrateVm详细参数
type PrimaryStorageMigrateVmDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"dstPrimaryStorageUuid" validate:"required"` // 必填
	rest string `json:"dstHostUuid,omitempty"`
	rest bool `json:"withDataVolumes,omitempty"`
	rest []string `json:"dataVolumeUuids,omitempty"`
	rest bool `json:"withSnapshots,omitempty"`
	rest int `json:"downTime,omitempty"`
	rest string `json:"strategy,omitempty"`
	rest int64 `json:"bandwidth,omitempty"`
	rest string `json:"volumeProvisioningStrategy,omitempty"`
}

// PrimaryStorageMigrateVmParam PrimaryStorageMigrateVm请求参数
type PrimaryStorageMigrateVmParam struct {
	BaseParam
	Params PrimaryStorageMigrateVmDetailParam `json:"params"` // 详细参数
}

