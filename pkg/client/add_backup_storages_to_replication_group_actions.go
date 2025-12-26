// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddBackupStoragesToReplicationGroup adds BackupStoragesToReplicationGroup
func (cli *ZSClient) AddBackupStoragesToReplicationGroup(params param.AddBackupStoragesToReplicationGroupParam) (*view.AddBackupStoragesToReplicationGroupEventView, error) {
	resp := view.AddBackupStoragesToReplicationGroupEventView{}
	if err := cli.Post("v1/image-replication-groups/{replicationGroupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
