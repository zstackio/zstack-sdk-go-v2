// Copyright (c) ZStack.io, Inc.

package param

// DownloadBackupFileFromPublicCloudDetailParam DownloadBackupFileFromPublicCloud详细参数
type DownloadBackupFileFromPublicCloudDetailParam struct {
	rest string `json:"regionId" validate:"required"` // 必填
	rest string `json:"file" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
}

// DownloadBackupFileFromPublicCloudParam DownloadBackupFileFromPublicCloud请求参数
type DownloadBackupFileFromPublicCloudParam struct {
	BaseParam
	Params DownloadBackupFileFromPublicCloudDetailParam `json:"params"` // 详细参数
}

