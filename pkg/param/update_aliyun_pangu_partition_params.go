// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunPanguPartitionDetailParam UpdateAliyunPanguPartition detail param
type UpdateAliyunPanguPartitionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AppName string `json:"appName,omitempty"`
	PartitionName string `json:"partitionName,omitempty"`
}

// UpdateAliyunPanguPartitionParam UpdateAliyunPanguPartition request param
type UpdateAliyunPanguPartitionParam struct {
	BaseParam
	Params UpdateAliyunPanguPartitionDetailParam `json:"params"`
}
