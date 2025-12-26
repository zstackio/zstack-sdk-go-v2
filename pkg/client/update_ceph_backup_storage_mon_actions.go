// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateCephBackupStorageMon updates CephBackupStorageMon
func (cli *ZSClient) UpdateCephBackupStorageMon(uuid string, params param.UpdateCephBackupStorageMonParam) (*view.UpdateCephBackupStorageMonEventView, error) {
	resp := view.UpdateCephBackupStorageMonEventView{}
	if err := cli.Put("v1/backup-storage/ceph/mons/{monUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
