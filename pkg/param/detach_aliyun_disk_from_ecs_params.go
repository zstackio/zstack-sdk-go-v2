// Copyright (c) ZStack.io, Inc.

package param

// DetachAliyunDiskFromEcsDetailParam DetachAliyunDiskFromEcs detail param
type DetachAliyunDiskFromEcsDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// DetachAliyunDiskFromEcsParam DetachAliyunDiskFromEcs request param
type DetachAliyunDiskFromEcsParam struct {
	BaseParam
	Params DetachAliyunDiskFromEcsDetailParam `json:"params"`
}
