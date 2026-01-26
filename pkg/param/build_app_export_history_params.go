// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteBuildAppExportHistoryParamDetail DeleteBuildAppExportHistory detail param
type DeleteBuildAppExportHistoryParamDetail struct {
	ExportId *string `json:"exportId,omitempty"`
	BuildSystemUuid *string `json:"buildSystemUuid,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteBuildAppExportHistoryParam DeleteBuildAppExportHistory request param
type DeleteBuildAppExportHistoryParam struct {
	BaseParam
	Params DeleteBuildAppExportHistoryParamDetail `json:"deleteBuildAppExportHistory"`
}
