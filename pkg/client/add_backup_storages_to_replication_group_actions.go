// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddBackupStoragesToReplicationGroup adds BackupStoragesToReplicationGroup
func (cli *ZSClient) AddBackupStoragesToReplicationGroup(params param.AddBackupStoragesToReplicationGroupParam) (*view.AddBackupStoragesToReplicationGroupEventView, error) {
	resp := view.AddBackupStoragesToReplicationGroupEventView{}
	if err := cli.Post("v1/image-replication-groups/{replicationGroupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
