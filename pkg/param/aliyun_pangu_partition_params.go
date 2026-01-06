// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddAliyunPanguPartitionParamDetail AddAliyunPanguPartition detail param
type AddAliyunPanguPartitionParamDetail struct {
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
	Params AddAliyunPanguPartitionParamDetail `json:"params"`
}
// DeleteAliyunPanguPartitionParamDetail DeleteAliyunPanguPartition detail param
type DeleteAliyunPanguPartitionParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunPanguPartitionParam DeleteAliyunPanguPartition request param
type DeleteAliyunPanguPartitionParam struct {
	BaseParam
	Params DeleteAliyunPanguPartitionParamDetail `json:"params"`
}
// UpdateAliyunPanguPartitionParamDetail UpdateAliyunPanguPartition detail param
type UpdateAliyunPanguPartitionParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AppName string `json:"appName,omitempty"`
	PartitionName string `json:"partitionName,omitempty"`
}

// UpdateAliyunPanguPartitionParam UpdateAliyunPanguPartition request param
type UpdateAliyunPanguPartitionParam struct {
	BaseParam
	Params UpdateAliyunPanguPartitionParamDetail `json:"params"`
}
