// Copyright (c) ZStack.io, Inc.

package param

// DeleteCephPrimaryStoragePoolDetailParam DeleteCephPrimaryStoragePool详细参数
type DeleteCephPrimaryStoragePoolDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// DeleteCephPrimaryStoragePoolParam DeleteCephPrimaryStoragePool请求参数
type DeleteCephPrimaryStoragePoolParam struct {
	BaseParam
	Params DeleteCephPrimaryStoragePoolDetailParam `json:"params"` // 详细参数
}

