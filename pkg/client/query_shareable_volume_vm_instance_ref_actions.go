// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryShareableVolumeVmInstanceRef queries ShareableVolumeVmInstanceRef list
func (cli *ZSClient) QueryShareableVolumeVmInstanceRef(params param.QueryParam) ([]view.ShareableVolumeVmInstanceRefInventoryView, error) {
	var resp []view.ShareableVolumeVmInstanceRefInventoryView
	return resp, cli.List("v1/volumes/vm-instances/refs", &params, &resp)
}
