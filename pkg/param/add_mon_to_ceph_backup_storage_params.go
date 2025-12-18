// Copyright (c) ZStack.io, Inc.

package param

// AddMonToCephBackupStorageDetailParam AddMonToCephBackupStorage详细参数
type AddMonToCephBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"monUrls" validate:"required"` // 必填
}

// AddMonToCephBackupStorageParam AddMonToCephBackupStorage请求参数
type AddMonToCephBackupStorageParam struct {
	BaseParam
	Params AddMonToCephBackupStorageDetailParam `json:"params"` // 详细参数
}

