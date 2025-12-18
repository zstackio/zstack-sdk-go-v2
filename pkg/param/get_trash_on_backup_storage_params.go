// Copyright (c) ZStack.io, Inc.

package param

// GetTrashOnBackupStorageDetailParam GetTrashOnBackupStorage detail param
type GetTrashOnBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	TrashType string `json:"trashType,omitempty"`
}

// GetTrashOnBackupStorageParam GetTrashOnBackupStorage request param
type GetTrashOnBackupStorageParam struct {
	BaseParam
	Params GetTrashOnBackupStorageDetailParam `json:"params"`
}
