// Copyright (c) ZStack.io, Inc.

package param

// AddMonToCephBackupStorageDetailParam AddMonToCephBackupStorage detail param
type AddMonToCephBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	MonUrls []string `json:"monUrls" validate:"required"`
}

// AddMonToCephBackupStorageParam AddMonToCephBackupStorage request param
type AddMonToCephBackupStorageParam struct {
	BaseParam
	Params AddMonToCephBackupStorageDetailParam `json:"params"`
}
