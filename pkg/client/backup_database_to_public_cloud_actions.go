// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// BackupDatabaseToPublicCloud 操作BackupDatabaseToPublicCloud
func (cli *ZSClient) BackupDatabaseToPublicCloud(params param.BackupDatabaseToPublicCloudParam) (*view.BackupDatabaseToPublicCloudEventView, error) {
	resp := view.BackupDatabaseToPublicCloudEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

