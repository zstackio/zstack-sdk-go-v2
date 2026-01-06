// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteDiskOfferingParamDetail DeleteDiskOffering detail param
type DeleteDiskOfferingParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteDiskOfferingParam DeleteDiskOffering request param
type DeleteDiskOfferingParam struct {
	BaseParam
	Params DeleteDiskOfferingParamDetail `json:"params"`
}
// CreateDiskOfferingParamDetail CreateDiskOffering detail param
type CreateDiskOfferingParamDetail struct {
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
	Params CreateDiskOfferingParamDetail `json:"params"`
}
// UpdateDiskOfferingParamDetail UpdateDiskOffering detail param
type UpdateDiskOfferingParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateDiskOfferingParam UpdateDiskOffering request param
type UpdateDiskOfferingParam struct {
	BaseParam
	Params UpdateDiskOfferingParamDetail `json:"params"`
}
