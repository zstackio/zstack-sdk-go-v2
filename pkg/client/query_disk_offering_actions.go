// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryDiskOffering queries DiskOffering list
func (cli *ZSClient) QueryDiskOffering(params param.QueryParam) ([]view.DiskOfferingInventoryView, error) {
	var resp []view.DiskOfferingInventoryView
	return resp, cli.List("v1/disk-offerings", &params, &resp)
}
