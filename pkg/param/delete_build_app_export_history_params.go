// Copyright (c) ZStack.io, Inc.

package param

// DeleteBuildAppExportHistoryDetailParam DeleteBuildAppExportHistory detail param
type DeleteBuildAppExportHistoryDetailParam struct {
	BuildAppUuid string `json:"buildAppUuid,omitempty"`
	ExportId string `json:"exportId,omitempty"`
	BuildSystemUuid string `json:"buildSystemUuid,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBuildAppExportHistoryParam DeleteBuildAppExportHistory request param
type DeleteBuildAppExportHistoryParam struct {
	BaseParam
	Params DeleteBuildAppExportHistoryDetailParam `json:"params"`
}
