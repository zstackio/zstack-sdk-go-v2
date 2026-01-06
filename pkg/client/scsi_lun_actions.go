// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
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
