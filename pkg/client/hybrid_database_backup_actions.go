// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// DeleteBackupFileInPublic deletes BackupFileInPublic
func (cli *ZSClient) DeleteBackupFileInPublic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/backup-mysql", uuid, string(deleteMode))
}

// BackupDatabaseToPublicCloud operates on DatabaseToPublicCloud
func (cli *ZSClient) BackupDatabaseToPublicCloud(params param.BackupDatabaseToPublicCloudParam) (*view.BackupDatabaseToPublicCloudEventView, error) {
	resp := view.BackupDatabaseToPublicCloudEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadBackupFileFromPublicCloud operates on DownloadBackupFileFromPublicCloud
func (cli *ZSClient) DownloadBackupFileFromPublicCloud(params param.DownloadBackupFileFromPublicCloudParam) (*view.DownloadBackupFileFromPublicCloudEventView, error) {
	resp := view.DownloadBackupFileFromPublicCloudEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql/download", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOssBackupBucketRemote creates OssBackupBucketRemote
func (cli *ZSClient) CreateOssBackupBucketRemote(params param.CreateOssBackupBucketRemoteParam) (*view.CreateOssBackupBucketRemoteEventView, error) {
	resp := view.CreateOssBackupBucketRemoteEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql/oss", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOssBackupBucketFromRemote gets OssBackupBucketFromRemote by uuid
func (cli *ZSClient) GetOssBackupBucketFromRemote() (*view.GetOssBackupBucketFromRemoteView, error) {
	var resp view.GetOssBackupBucketFromRemoteView
	if err := cli.GetWithRespKey("v1/hybrid/backup-mysql/oss", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
