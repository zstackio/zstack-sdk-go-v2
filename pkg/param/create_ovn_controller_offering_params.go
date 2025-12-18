// Copyright (c) ZStack.io, Inc.

package param

// CreateOvnControllerOfferingDetailParam CreateOvnControllerOffering detail param
type CreateOvnControllerOfferingDetailParam struct {
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

// CreateOvnControllerOfferingParam CreateOvnControllerOffering request param
type CreateOvnControllerOfferingParam struct {
	BaseParam
	Params CreateOvnControllerOfferingDetailParam `json:"params"`
}
