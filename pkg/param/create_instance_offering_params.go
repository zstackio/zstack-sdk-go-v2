// Copyright (c) ZStack.io, Inc.

package param

// CreateInstanceOfferingDetailParam CreateInstanceOffering detail param
type CreateInstanceOfferingDetailParam struct {
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

// CreateInstanceOfferingParam CreateInstanceOffering request param
type CreateInstanceOfferingParam struct {
	BaseParam
	Params CreateInstanceOfferingDetailParam `json:"params"`
}
