// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteExternalBackupParamDetail DeleteExternalBackup detail param
type DeleteExternalBackupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteExternalBackupParam DeleteExternalBackup request param
type DeleteExternalBackupParam struct {
	BaseParam
	DeleteExternalBackup DeleteExternalBackupParamDetail `json:"deleteExternalBackup"`
}
