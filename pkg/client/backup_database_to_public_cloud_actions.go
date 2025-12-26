// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// BackupDatabaseToPublicCloud operates on BackupDatabaseToPublicCloud
func (cli *ZSClient) BackupDatabaseToPublicCloud(params param.BackupDatabaseToPublicCloudParam) (*view.BackupDatabaseToPublicCloudEventView, error) {
	resp := view.BackupDatabaseToPublicCloudEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
