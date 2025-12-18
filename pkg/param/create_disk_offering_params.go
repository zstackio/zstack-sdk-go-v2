// Copyright (c) ZStack.io, Inc.

package param

// CreateDiskOfferingDetailParam CreateDiskOffering detail param
type CreateDiskOfferingDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	DiskSize int64 `json:"diskSize" validate:"required"`
	SortKey int `json:"sortKey,omitempty"`
	AllocationStrategy string `json:"allocationStrategy,omitempty"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDiskOfferingParam CreateDiskOffering request param
type CreateDiskOfferingParam struct {
	BaseParam
	Params CreateDiskOfferingDetailParam `json:"params"`
}
