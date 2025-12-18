// Copyright (c) ZStack.io, Inc.

package param

// ConvertVmFromForeignHypervisorDetailParam ConvertVmFromForeignHypervisor详细参数
type ConvertVmFromForeignHypervisorDetailParam struct {
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"conversionHostUuid,omitempty"`
	rest string `json:"sshPrivKey,omitempty"`
	rest int `json:"cpuNum" validate:"required"` // 必填
	rest int64 `json:"memorySize" validate:"required"` // 必填
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuid" validate:"required"` // 必填
	rest []string `json:"l3NetworkUuids" validate:"required"` // 必填
	rest string `json:"defaultL3NetworkUuid,omitempty"`
	rest string `json:"platform,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"strategy,omitempty"`
	rest string `json:"convertStrategy,omitempty"`
	rest bool `json:"pauseVm,omitempty"`
	rest []interface{} `json:"volumeFilters,omitempty"`
	rest string `json:"rootFileSystem,omitempty"`
	rest string `json:"longJobName,omitempty"`
	rest string `json:"longJobDescription,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// ConvertVmFromForeignHypervisorParam ConvertVmFromForeignHypervisor请求参数
type ConvertVmFromForeignHypervisorParam struct {
	BaseParam
	Params ConvertVmFromForeignHypervisorDetailParam `json:"params"` // 详细参数
}

