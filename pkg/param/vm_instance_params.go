// Copyright (c) ZStack.io, Inc.

package param

// CloneVmInstanceDetailParam CloneVmInstance详细参数
type CloneVmInstanceDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"strategy,omitempty"`
	rest []interface{} `json:"vmNicParams,omitempty"`
	rest []string `json:"names" validate:"required"` // 必填
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest string `json:"primaryStorageUuidForDataVolume,omitempty"`
	rest bool `json:"full,omitempty"`
	rest []string `json:"rootVolumeSystemTags,omitempty"`
	rest []string `json:"dataVolumeSystemTags,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
}

// CloneVmInstanceParam CloneVmInstance请求参数
type CloneVmInstanceParam struct {
	BaseParam
	Params CloneVmInstanceDetailParam `json:"params"` // 详细参数
}

