// Copyright (c) ZStack.io, Inc.

package param

// CreateVirtualRouterOfferingDetailParam CreateVirtualRouterOffering detail param
type CreateVirtualRouterOfferingDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ManagementNetworkUuid string `json:"managementNetworkUuid" validate:"required"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	PublicNetworkUuid string `json:"publicNetworkUuid,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
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

// CreateVirtualRouterOfferingParam CreateVirtualRouterOffering request param
type CreateVirtualRouterOfferingParam struct {
	BaseParam
	Params CreateVirtualRouterOfferingDetailParam `json:"params"`
}
