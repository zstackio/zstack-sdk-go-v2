// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// ChangeInstanceOfferingParamDetail ChangeInstanceOffering detail param
type ChangeInstanceOfferingParamDetail struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid" validate:"required"`
}

// ChangeInstanceOfferingParam ChangeInstanceOffering request param
type ChangeInstanceOfferingParam struct {
	BaseParam
	Params ChangeInstanceOfferingParamDetail `json:"params"`
}
// UpdateInstanceOfferingParamDetail UpdateInstanceOffering detail param
type UpdateInstanceOfferingParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
}

// UpdateInstanceOfferingParam UpdateInstanceOffering request param
type UpdateInstanceOfferingParam struct {
	BaseParam
	Params UpdateInstanceOfferingParamDetail `json:"params"`
}
// CreateInstanceOfferingParamDetail CreateInstanceOffering detail param
type CreateInstanceOfferingParamDetail struct {
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
	Params CreateInstanceOfferingParamDetail `json:"params"`
}
// DeleteInstanceOfferingParamDetail DeleteInstanceOffering detail param
type DeleteInstanceOfferingParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteInstanceOfferingParam DeleteInstanceOffering request param
type DeleteInstanceOfferingParam struct {
	BaseParam
	Params DeleteInstanceOfferingParamDetail `json:"params"`
}
