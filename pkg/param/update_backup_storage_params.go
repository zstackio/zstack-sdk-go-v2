// Copyright (c) ZStack.io, Inc.

package param

// UpdateBackupStorageDetailParam UpdateBackupStorage detail param
type UpdateBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateBackupStorageParam UpdateBackupStorage request param
type UpdateBackupStorageParam struct {
	BaseParam
	Params UpdateBackupStorageDetailParam `json:"params"`
}
