// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVCenterBackupStorage queries VCenterBackupStorage list
func (cli *ZSClient) QueryVCenterBackupStorage(params *param.QueryParam) ([]view.VCenterBackupStorageInventoryView, error) {
	var resp []view.VCenterBackupStorageInventoryView
	return resp, cli.List("v1/vcenters/backup-storage", params, &resp)
}

// PageVCenterBackupStorage Pagination
func (cli *ZSClient) PageVCenterBackupStorage(params *param.QueryParam) ([]view.VCenterBackupStorageInventoryView, int, error) {
	var vCenterBackupStorages []view.VCenterBackupStorageInventoryView
	total, err := cli.Page("v1/vcenters/backup-storage", params, &vCenterBackupStorages)
	return vCenterBackupStorages, total, err
}
