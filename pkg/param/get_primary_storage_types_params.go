// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageTypesDetailParam GetPrimaryStorageTypes detail param
type GetPrimaryStorageTypesDetailParam struct {
}

// GetPrimaryStorageTypesParam GetPrimaryStorageTypes request param
type GetPrimaryStorageTypesParam struct {
	BaseParam
	Params GetPrimaryStorageTypesDetailParam `json:"params"`
}
