// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddCephBackupStorage adds CephBackupStorage
func (cli *ZSClient) AddCephBackupStorage(params param.AddCephBackupStorageParam) (*view.AddBackupStorageEventView, error) {
	resp := view.AddBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/ceph", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
