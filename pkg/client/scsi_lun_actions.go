// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateScsiLun updates ScsiLun
func (cli *ZSClient) UpdateScsiLun(uuid string, params param.UpdateScsiLunParam) (*view.ScsiLunInventoryView, error) {
	var resp view.UpdateScsiLunEventView
	if err := cli.Put("v1/storage-devices/scsi-lun/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryScsiLun queries ScsiLun list
func (cli *ZSClient) QueryScsiLun(params *param.QueryParam) ([]view.ScsiLunInventoryView, error) {
	var resp []view.ScsiLunInventoryView
	return resp, cli.List("v1/storage-devices/scsi-lun/luns", params, &resp)
}

func (cli *ZSClient) GetScsiLun(uuid string) (*view.ScsiLunInventoryView, error) {
	var resp view.ScsiLunInventoryView
	if err := cli.Get("v1/storage-devices/scsi-lun/luns", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
