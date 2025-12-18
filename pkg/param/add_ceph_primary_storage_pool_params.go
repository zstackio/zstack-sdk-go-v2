// Copyright (c) ZStack.io, Inc.

package param

// AddCephPrimaryStoragePoolDetailParam AddCephPrimaryStoragePool详细参数
type AddCephPrimaryStoragePoolDetailParam struct {
	rest string `json:"primaryStorageUuid" validate:"required"` // 必填
	rest string `json:"poolName" validate:"required"` // 必填
	rest string `json:"aliasName,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type" validate:"required"` // 必填
	rest bool `json:"isCreate,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddCephPrimaryStoragePoolParam AddCephPrimaryStoragePool请求参数
type AddCephPrimaryStoragePoolParam struct {
	BaseParam
	Params AddCephPrimaryStoragePoolDetailParam `json:"params"` // 详细参数
}

