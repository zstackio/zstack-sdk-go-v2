// Copyright (c) ZStack.io, Inc.

package param

// RemoveMonFromCephBackupStorageDetailParam RemoveMonFromCephBackupStorage详细参数
type RemoveMonFromCephBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"monHostnames" validate:"required"` // 必填
}

// RemoveMonFromCephBackupStorageParam RemoveMonFromCephBackupStorage请求参数
type RemoveMonFromCephBackupStorageParam struct {
	BaseParam
	Params RemoveMonFromCephBackupStorageDetailParam `json:"params"` // 详细参数
}

