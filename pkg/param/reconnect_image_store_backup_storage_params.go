// Copyright (c) ZStack.io, Inc.

package param

// ReconnectImageStoreBackupStorageDetailParam ReconnectImageStoreBackupStorage detail param
type ReconnectImageStoreBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectImageStoreBackupStorageParam ReconnectImageStoreBackupStorage request param
type ReconnectImageStoreBackupStorageParam struct {
	BaseParam
	Params ReconnectImageStoreBackupStorageDetailParam `json:"params"`
}
