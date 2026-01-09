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

func (cli *ZSClient) GetVCenterBackupStorage(uuid string) (*view.VCenterBackupStorageInventoryView, error) {
	var resp view.VCenterBackupStorageInventoryView
	if err := cli.Get("v1/vcenters/backup-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
