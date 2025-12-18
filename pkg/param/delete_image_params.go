// Copyright (c) ZStack.io, Inc.

package param

// DeleteImageDetailParam DeleteImage detail param
type DeleteImageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteImageParam DeleteImage request param
type DeleteImageParam struct {
	BaseParam
	Params DeleteImageDetailParam `json:"params"`
}
