// Copyright (c) ZStack.io, Inc.

package param

// DeleteExternalBackupDetailParam DeleteExternalBackup detail param
type DeleteExternalBackupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteExternalBackupParam DeleteExternalBackup request param
type DeleteExternalBackupParam struct {
	BaseParam
	Params DeleteExternalBackupDetailParam `json:"params"`
}
