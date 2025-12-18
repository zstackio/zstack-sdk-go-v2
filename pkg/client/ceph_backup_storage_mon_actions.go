// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateCephBackupStorageMon 更新CephBackupStorageMon
func (cli *ZSClient) UpdateCephBackupStorageMon(uuid string, params param.UpdateCephBackupStorageMonParam) (*view.UpdateCephBackupStorageMonEventView, error) {
	resp := view.UpdateCephBackupStorageMonEventView{}
	if err := cli.Put("v1/backup-storage/ceph/mons/{monUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

