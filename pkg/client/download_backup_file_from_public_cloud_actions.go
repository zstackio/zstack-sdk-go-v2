// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DownloadBackupFileFromPublicCloud operates on DownloadBackupFileFromPublicCloud
func (cli *ZSClient) DownloadBackupFileFromPublicCloud(params param.DownloadBackupFileFromPublicCloudParam) (*view.DownloadBackupFileFromPublicCloudEventView, error) {
	resp := view.DownloadBackupFileFromPublicCloudEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql/download", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
