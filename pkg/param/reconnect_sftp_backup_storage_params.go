// Copyright (c) ZStack.io, Inc.

package param

// ReconnectSftpBackupStorageDetailParam ReconnectSftpBackupStorage详细参数
type ReconnectSftpBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectSftpBackupStorageParam ReconnectSftpBackupStorage请求参数
type ReconnectSftpBackupStorageParam struct {
	BaseParam
	Params ReconnectSftpBackupStorageDetailParam `json:"params"` // 详细参数
}

