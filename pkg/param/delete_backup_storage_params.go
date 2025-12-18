// Copyright (c) ZStack.io, Inc.

package param

// DeleteBackupStorageDetailParam DeleteBackupStorage detail param
type DeleteBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBackupStorageParam DeleteBackupStorage request param
type DeleteBackupStorageParam struct {
	BaseParam
	Params DeleteBackupStorageDetailParam `json:"params"`
}
