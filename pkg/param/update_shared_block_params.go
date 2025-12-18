// Copyright (c) ZStack.io, Inc.

package param

// UpdateSharedBlockDetailParam UpdateSharedBlock detail param
type UpdateSharedBlockDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SharedBlockGroupUuid string `json:"sharedBlockGroupUuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DiskUuid string `json:"diskUuid,omitempty"`
}

// UpdateSharedBlockParam UpdateSharedBlock request param
type UpdateSharedBlockParam struct {
	BaseParam
	Params UpdateSharedBlockDetailParam `json:"params"`
}
