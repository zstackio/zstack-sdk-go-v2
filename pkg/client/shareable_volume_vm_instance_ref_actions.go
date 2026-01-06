// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryShareableVolumeVmInstanceRef queries ShareableVolumeVmInstanceRef list
func (cli *ZSClient) QueryShareableVolumeVmInstanceRef(params *param.QueryParam) ([]view.ShareableVolumeVmInstanceRefInventoryView, error) {
	var resp []view.ShareableVolumeVmInstanceRefInventoryView
	return resp, cli.List("v1/volumes/vm-instances/refs", params, &resp)
}
