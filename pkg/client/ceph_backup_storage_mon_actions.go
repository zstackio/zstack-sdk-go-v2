// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateCephBackupStorageMon updates CephBackupStorageMon
func (cli *ZSClient) UpdateCephBackupStorageMon(monUuid string, params param.UpdateCephBackupStorageMonParam) (*view.CephBackupStorageInventoryView, error) {
	var resp view.UpdateCephBackupStorageMonEventView
	err := cli.PutWithSpec("v1/backup-storage/ceph/mons", fmt.Sprintf(\"%s/actions\", monUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
