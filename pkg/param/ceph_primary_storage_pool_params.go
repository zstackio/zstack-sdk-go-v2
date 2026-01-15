// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateCephPrimaryStoragePoolParamDetail UpdateCephPrimaryStoragePool detail param
type UpdateCephPrimaryStoragePoolParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	AliasName string `json:"aliasName,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateCephPrimaryStoragePoolParam UpdateCephPrimaryStoragePool request param
type UpdateCephPrimaryStoragePoolParam struct {
	BaseParam
	Params UpdateCephPrimaryStoragePoolParamDetail `json:"updateCephPrimaryStoragePool"`
}
// AddCephPrimaryStoragePoolParamDetail AddCephPrimaryStoragePool detail param
type AddCephPrimaryStoragePoolParamDetail struct {
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
	PoolName string `json:"poolName" validate:"required"`
	AliasName string `json:"aliasName,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	IsCreate bool `json:"isCreate,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddCephPrimaryStoragePoolParam AddCephPrimaryStoragePool request param
type AddCephPrimaryStoragePoolParam struct {
	BaseParam
	Params AddCephPrimaryStoragePoolParamDetail `json:"addCephPrimaryStoragePool"`
}
// DeleteCephPrimaryStoragePoolParamDetail DeleteCephPrimaryStoragePool detail param
type DeleteCephPrimaryStoragePoolParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteCephPrimaryStoragePoolParam DeleteCephPrimaryStoragePool request param
type DeleteCephPrimaryStoragePoolParam struct {
	BaseParam
	Params DeleteCephPrimaryStoragePoolParamDetail `json:"deleteCephPrimaryStoragePool"`
}
