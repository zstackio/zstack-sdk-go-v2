// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunPanguPartitionDetailParam AddAliyunPanguPartition detail param
type AddAliyunPanguPartitionDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	IdentityZoneUuid string `json:"identityZoneUuid" validate:"required"`
	AppName string `json:"appName" validate:"required"`
	PartitionName string `json:"partitionName" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunPanguPartitionParam AddAliyunPanguPartition request param
type AddAliyunPanguPartitionParam struct {
	BaseParam
	Params AddAliyunPanguPartitionDetailParam `json:"params"`
}
