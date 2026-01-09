// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateCephBackupStorageMon updates CephBackupStorageMon
func (cli *ZSClient) UpdateCephBackupStorageMon(uuid string, params param.UpdateCephBackupStorageMonParam) (*view.CephBackupStorageInventoryView, error) {
	var resp view.UpdateCephBackupStorageMonEventView
	if err := cli.Put("v1/backup-storage/ceph/mons/{monUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
