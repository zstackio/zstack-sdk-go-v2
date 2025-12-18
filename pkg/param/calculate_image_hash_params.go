// Copyright (c) ZStack.io, Inc.

package param

// CalculateImageHashDetailParam CalculateImageHash detail param
type CalculateImageHashDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	Algorithm string `json:"algorithm,omitempty"`
}

// CalculateImageHashParam CalculateImageHash request param
type CalculateImageHashParam struct {
	BaseParam
	Params CalculateImageHashDetailParam `json:"params"`
}
