// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// DeleteBackupFileInPublic requires additional arguments.
//
// Deprecated: use DeleteBackupFileInPublicWithParams to supply the required fields.
func (cli *ZSClient) DeleteBackupFileInPublic(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DeleteBackupFileInPublic requires additional arguments; use DeleteBackupFileInPublicWithParams")
}

// DeleteBackupFileInPublicWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteBackupFileInPublicWithParams(deleteMode param.DeleteMode, params param.DeleteBackupFileInPublicParamDetail) error {
	query := url.Values{
		"type":     {params.Type},
		"regionId": {params.RegionId},
		"file":     {params.File},
	}
	return cli.deleteActionWithQuery("v1/hybrid/backup-mysql", "", "", deleteModeOrDefault(deleteMode, params.DeleteMode), query, "type", "regionId", "file")
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
