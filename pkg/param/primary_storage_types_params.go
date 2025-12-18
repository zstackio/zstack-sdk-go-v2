// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageTypesDetailParam GetPrimaryStorageTypes详细参数
type GetPrimaryStorageTypesDetailParam struct {
}

// GetPrimaryStorageTypesParam GetPrimaryStorageTypes请求参数
type GetPrimaryStorageTypesParam struct {
	BaseParam
	Params GetPrimaryStorageTypesDetailParam `json:"params"` // 详细参数
}

