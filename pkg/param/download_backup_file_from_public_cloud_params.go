// Copyright (c) ZStack.io, Inc.

package param

// DownloadBackupFileFromPublicCloudDetailParam DownloadBackupFileFromPublicCloud detail param
type DownloadBackupFileFromPublicCloudDetailParam struct {
	RegionId string `json:"regionId" validate:"required"`
	File string `json:"file" validate:"required"`
	Type string `json:"type" validate:"required"`
}

// DownloadBackupFileFromPublicCloudParam DownloadBackupFileFromPublicCloud request param
type DownloadBackupFileFromPublicCloudParam struct {
	BaseParam
	Params DownloadBackupFileFromPublicCloudDetailParam `json:"params"`
}
