// Copyright (c) ZStack.io, Inc.

package param

// DetachAliyunDiskFromEcsDetailParam DetachAliyunDiskFromEcs详细参数
type DetachAliyunDiskFromEcsDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// DetachAliyunDiskFromEcsParam DetachAliyunDiskFromEcs请求参数
type DetachAliyunDiskFromEcsParam struct {
	BaseParam
	Params DetachAliyunDiskFromEcsDetailParam `json:"params"` // 详细参数
}

