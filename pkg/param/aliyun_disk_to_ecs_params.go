// Copyright (c) ZStack.io, Inc.

package param

// AttachAliyunDiskToEcsDetailParam AttachAliyunDiskToEcs详细参数
type AttachAliyunDiskToEcsDetailParam struct {
	rest string `json:"ecsUuid" validate:"required"` // 必填
	rest string `json:"diskUuid" validate:"required"` // 必填
	rest bool `json:"deleteWithInstance,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AttachAliyunDiskToEcsParam AttachAliyunDiskToEcs请求参数
type AttachAliyunDiskToEcsParam struct {
	BaseParam
	Params AttachAliyunDiskToEcsDetailParam `json:"params"` // 详细参数
}

