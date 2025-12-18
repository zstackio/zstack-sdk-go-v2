// Copyright (c) ZStack.io, Inc.

package param

// CreateFaultToleranceVmInstanceDetailParam CreateFaultToleranceVmInstance详细参数
type CreateFaultToleranceVmInstanceDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"instanceOfferingUuid,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest int64 `json:"reservedMemorySize,omitempty"`
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest []string `json:"l3NetworkUuids" validate:"required"` // 必填
	rest []string `json:"dataDiskOfferingUuids,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"defaultL3NetworkUuid,omitempty"`
	rest []string `json:"rootVolumeSystemTags,omitempty"`
	rest []string `json:"dataVolumeSystemTags,omitempty"`
	rest string `json:"strategy,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateFaultToleranceVmInstanceParam CreateFaultToleranceVmInstance请求参数
type CreateFaultToleranceVmInstanceParam struct {
	BaseParam
	Params CreateFaultToleranceVmInstanceDetailParam `json:"params"` // 详细参数
}

