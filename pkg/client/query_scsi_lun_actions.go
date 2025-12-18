// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryScsiLun queries ScsiLun list
func (cli *ZSClient) QueryScsiLun(params param.QueryParam) ([]view.ScsiLunInventoryView, error) {
	var resp []view.ScsiLunInventoryView
	return resp, cli.List("v1/storage-devices/scsi-lun/luns", &params, &resp)
}
