// Copyright (c) ZStack.io, Inc.

package param

// CreateOvnControllerVmDetailParam CreateOvnControllerVm详细参数
type CreateOvnControllerVmDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"instanceOfferingUuid,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest int64 `json:"reservedMemorySize,omitempty"`
	rest string `json:"imageUuid,omitempty"`
	rest []string `json:"l3NetworkUuids,omitempty"`
	rest string `json:"vmNicParams,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"rootDiskOfferingUuid,omitempty"`
	rest int64 `json:"rootDiskSize,omitempty"`
	rest []int64 `json:"dataDiskSizes,omitempty"`
	rest []string `json:"dataDiskOfferingUuids,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"defaultL3NetworkUuid,omitempty"`
	rest string `json:"strategy,omitempty"`
	rest []string `json:"rootVolumeSystemTags,omitempty"`
	rest []string `json:"dataVolumeSystemTags,omitempty"`
	rest map[string]interface{} `json:"dataVolumeSystemTagsOnIndex,omitempty"`
	rest []string `json:"sshKeyPairUuids,omitempty"`
	rest string `json:"platform,omitempty"`
	rest string `json:"guestOsType,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest bool `json:"virtio,omitempty"`
	rest []interface{} `json:"diskAOs,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateOvnControllerVmParam CreateOvnControllerVm请求参数
type CreateOvnControllerVmParam struct {
	BaseParam
	Params CreateOvnControllerVmDetailParam `json:"params"` // 详细参数
}

