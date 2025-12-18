// Copyright (c) ZStack.io, Inc.

package param

// AddSftpBackupStorageDetailParam AddSftpBackupStorage详细参数
type AddSftpBackupStorageDetailParam struct {
	rest string `json:"hostname" validate:"required"` // 必填
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest int `json:"sshPort,omitempty"`
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest bool `json:"importImages,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSftpBackupStorageParam AddSftpBackupStorage请求参数
type AddSftpBackupStorageParam struct {
	BaseParam
	Params AddSftpBackupStorageDetailParam `json:"params"` // 详细参数
}

