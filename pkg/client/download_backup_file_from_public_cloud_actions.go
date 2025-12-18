// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DownloadBackupFileFromPublicCloud 操作DownloadBackupFileFromPublicCloud
func (cli *ZSClient) DownloadBackupFileFromPublicCloud(params param.DownloadBackupFileFromPublicCloudParam) (*view.DownloadBackupFileFromPublicCloudEventView, error) {
	resp := view.DownloadBackupFileFromPublicCloudEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql/download", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

