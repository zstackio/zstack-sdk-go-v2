// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateNfvInstOfferingParamDetail CreateNfvInstOffering detail param
type CreateNfvInstOfferingParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ManagementNetworkUuid string `json:"managementNetworkUuid" validate:"required"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	CpuNum int `json:"cpuNum" validate:"required"`
	MemorySize int64 `json:"memorySize" validate:"required"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
	SortKey int `json:"sortKey,omitempty"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateNfvInstOfferingParam CreateNfvInstOffering request param
type CreateNfvInstOfferingParam struct {
	BaseParam
	Params CreateNfvInstOfferingParamDetail `json:"createNfvInstOffering"`
}
