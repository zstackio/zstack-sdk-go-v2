// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
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
