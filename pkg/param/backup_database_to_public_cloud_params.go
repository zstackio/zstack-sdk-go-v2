// Copyright (c) ZStack.io, Inc.

package param

// BackupDatabaseToPublicCloudDetailParam BackupDatabaseToPublicCloud detail param
type BackupDatabaseToPublicCloudDetailParam struct {
	Type string `json:"type" validate:"required"`
	RegionId string `json:"regionId" validate:"required"`
	Local bool `json:"local,omitempty"`
}

// BackupDatabaseToPublicCloudParam BackupDatabaseToPublicCloud request param
type BackupDatabaseToPublicCloudParam struct {
	BaseParam
	Params BackupDatabaseToPublicCloudDetailParam `json:"params"`
}
