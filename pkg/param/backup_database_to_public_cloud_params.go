// Copyright (c) ZStack.io, Inc.

package param

// BackupDatabaseToPublicCloudDetailParam BackupDatabaseToPublicCloud详细参数
type BackupDatabaseToPublicCloudDetailParam struct {
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"regionId" validate:"required"` // 必填
	rest bool `json:"local,omitempty"`
}

// BackupDatabaseToPublicCloudParam BackupDatabaseToPublicCloud请求参数
type BackupDatabaseToPublicCloudParam struct {
	BaseParam
	Params BackupDatabaseToPublicCloudDetailParam `json:"params"` // 详细参数
}

