// Copyright (c) ZStack.io, Inc.

package param

// CleanUpTrashOnPrimaryStorageDetailParam CleanUpTrashOnPrimaryStorage detail param
type CleanUpTrashOnPrimaryStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	TrashId int64 `json:"trashId,omitempty"`
}

// CleanUpTrashOnPrimaryStorageParam CleanUpTrashOnPrimaryStorage request param
type CleanUpTrashOnPrimaryStorageParam struct {
	BaseParam
	Params CleanUpTrashOnPrimaryStorageDetailParam `json:"params"`
}
