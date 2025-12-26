// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryDiskOffering queries DiskOffering list
func (cli *ZSClient) QueryDiskOffering(params *param.QueryParam) ([]view.DiskOfferingInventoryView, error) {
	var resp []view.DiskOfferingInventoryView
	return resp, cli.List("v1/disk-offerings", params, &resp)
}
