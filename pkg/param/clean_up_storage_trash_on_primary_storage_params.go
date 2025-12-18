// Copyright (c) ZStack.io, Inc.

package param

// CleanUpStorageTrashOnPrimaryStorageDetailParam CleanUpStorageTrashOnPrimaryStorage detail param
type CleanUpStorageTrashOnPrimaryStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Force bool `json:"force,omitempty"`
}

// CleanUpStorageTrashOnPrimaryStorageParam CleanUpStorageTrashOnPrimaryStorage request param
type CleanUpStorageTrashOnPrimaryStorageParam struct {
	BaseParam
	Params CleanUpStorageTrashOnPrimaryStorageDetailParam `json:"params"`
}
