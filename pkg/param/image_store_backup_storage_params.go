// Copyright (c) ZStack.io, Inc.

package param

// UpdateImageStoreBackupStorageDetailParam UpdateImageStoreBackupStorage详细参数
type UpdateImageStoreBackupStorageDetailParam struct {
	rest string `json:"username,omitempty"`
	rest string `json:"password,omitempty"`
	rest string `json:"hostname,omitempty"`
	rest int `json:"sshPort,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateImageStoreBackupStorageParam UpdateImageStoreBackupStorage请求参数
type UpdateImageStoreBackupStorageParam struct {
	BaseParam
	Params UpdateImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

