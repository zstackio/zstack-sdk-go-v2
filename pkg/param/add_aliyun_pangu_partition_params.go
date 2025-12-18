// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunPanguPartitionDetailParam AddAliyunPanguPartition详细参数
type AddAliyunPanguPartitionDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"identityZoneUuid" validate:"required"` // 必填
	rest string `json:"appName" validate:"required"` // 必填
	rest string `json:"partitionName" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAliyunPanguPartitionParam AddAliyunPanguPartition请求参数
type AddAliyunPanguPartitionParam struct {
	BaseParam
	Params AddAliyunPanguPartitionDetailParam `json:"params"` // 详细参数
}

