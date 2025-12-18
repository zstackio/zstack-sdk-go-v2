// Copyright (c) ZStack.io, Inc.

package param

// GetTrashOnPrimaryStorageDetailParam GetTrashOnPrimaryStorage detail param
type GetTrashOnPrimaryStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	TrashType string `json:"trashType,omitempty"`
}

// GetTrashOnPrimaryStorageParam GetTrashOnPrimaryStorage request param
type GetTrashOnPrimaryStorageParam struct {
	BaseParam
	Params GetTrashOnPrimaryStorageDetailParam `json:"params"`
}
