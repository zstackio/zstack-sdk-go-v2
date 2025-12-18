// Copyright (c) ZStack.io, Inc.

package param

// AttachAliyunDiskToEcsDetailParam AttachAliyunDiskToEcs detail param
type AttachAliyunDiskToEcsDetailParam struct {
	EcsUuid string `json:"ecsUuid" validate:"required"`
	DiskUuid string `json:"diskUuid" validate:"required"`
	DeleteWithInstance bool `json:"deleteWithInstance,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AttachAliyunDiskToEcsParam AttachAliyunDiskToEcs request param
type AttachAliyunDiskToEcsParam struct {
	BaseParam
	Params AttachAliyunDiskToEcsDetailParam `json:"params"`
}
