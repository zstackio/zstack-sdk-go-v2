// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryScsiLun queries ScsiLun list
func (cli *ZSClient) QueryScsiLun(params *param.QueryParam) ([]view.ScsiLunInventoryView, error) {
	var resp []view.ScsiLunInventoryView
	return resp, cli.List("v1/storage-devices/scsi-lun/luns", params, &resp)
}
