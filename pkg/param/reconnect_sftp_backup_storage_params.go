// Copyright (c) ZStack.io, Inc.

package param

// ReconnectSftpBackupStorageDetailParam ReconnectSftpBackupStorage detail param
type ReconnectSftpBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectSftpBackupStorageParam ReconnectSftpBackupStorage request param
type ReconnectSftpBackupStorageParam struct {
	BaseParam
	Params ReconnectSftpBackupStorageDetailParam `json:"params"`
}
