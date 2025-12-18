// Copyright (c) ZStack.io, Inc.

package param

// AddCephPrimaryStoragePoolDetailParam AddCephPrimaryStoragePool detail param
type AddCephPrimaryStoragePoolDetailParam struct {
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
	Params AddCephPrimaryStoragePoolDetailParam `json:"params"`
}
