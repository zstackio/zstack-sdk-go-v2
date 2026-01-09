// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddExternalBackupStorage adds ExternalBackupStorage
func (cli *ZSClient) AddExternalBackupStorage(params param.AddExternalBackupStorageParam) (*view.ExternalBackupStorageInventoryView, error) {
	var resp view.AddExternalBackupStorageEventView
	if err := cli.Post("v1/backup-storage/addon", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
