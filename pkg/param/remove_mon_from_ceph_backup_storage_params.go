// Copyright (c) ZStack.io, Inc.

package param

// RemoveMonFromCephBackupStorageDetailParam RemoveMonFromCephBackupStorage detail param
type RemoveMonFromCephBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	MonHostnames []string `json:"monHostnames" validate:"required"`
}

// RemoveMonFromCephBackupStorageParam RemoveMonFromCephBackupStorage request param
type RemoveMonFromCephBackupStorageParam struct {
	BaseParam
	Params RemoveMonFromCephBackupStorageDetailParam `json:"params"`
}
