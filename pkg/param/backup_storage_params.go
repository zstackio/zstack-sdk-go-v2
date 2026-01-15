// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateBackupStorageParamDetail UpdateBackupStorage detail param
type UpdateBackupStorageParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateBackupStorageParam UpdateBackupStorage request param
type UpdateBackupStorageParam struct {
	BaseParam
	UpdateBackupStorage UpdateBackupStorageParamDetail `json:"updateBackupStorage"`
}
// DeleteBackupStorageParamDetail DeleteBackupStorage detail param
type DeleteBackupStorageParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBackupStorageParam DeleteBackupStorage request param
type DeleteBackupStorageParam struct {
	BaseParam
	DeleteBackupStorage DeleteBackupStorageParamDetail `json:"deleteBackupStorage"`
}
// ReconnectBackupStorageParamDetail ReconnectBackupStorage detail param
type ReconnectBackupStorageParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectBackupStorageParam ReconnectBackupStorage request param
type ReconnectBackupStorageParam struct {
	BaseParam
	ReconnectBackupStorage ReconnectBackupStorageParamDetail `json:"reconnectBackupStorage"`
}
