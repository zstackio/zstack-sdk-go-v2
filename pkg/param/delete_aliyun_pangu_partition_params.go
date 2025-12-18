// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunPanguPartitionDetailParam DeleteAliyunPanguPartition detail param
type DeleteAliyunPanguPartitionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunPanguPartitionParam DeleteAliyunPanguPartition request param
type DeleteAliyunPanguPartitionParam struct {
	BaseParam
	Params DeleteAliyunPanguPartitionDetailParam `json:"params"`
}
