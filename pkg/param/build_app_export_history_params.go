// Copyright (c) ZStack.io, Inc.

package param

// DeleteBuildAppExportHistoryDetailParam DeleteBuildAppExportHistory详细参数
type DeleteBuildAppExportHistoryDetailParam struct {
	rest string `json:"buildAppUuid,omitempty"`
	rest string `json:"exportId,omitempty"`
	rest string `json:"buildSystemUuid,omitempty"`
	rest string `json:"deleteMode,omitempty"`
}

// DeleteBuildAppExportHistoryParam DeleteBuildAppExportHistory请求参数
type DeleteBuildAppExportHistoryParam struct {
	BaseParam
	Params DeleteBuildAppExportHistoryDetailParam `json:"params"` // 详细参数
}

