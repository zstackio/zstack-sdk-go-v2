// Copyright (c) ZStack.io, Inc.

package param

// ReconnectBackupStorageDetailParam ReconnectBackupStorage detail param
type ReconnectBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectBackupStorageParam ReconnectBackupStorage request param
type ReconnectBackupStorageParam struct {
	BaseParam
	Params ReconnectBackupStorageDetailParam `json:"params"`
}
