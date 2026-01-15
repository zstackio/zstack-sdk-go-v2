// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryShareableVolumeVmInstanceRef queries ShareableVolumeVmInstanceRef list
func (cli *ZSClient) QueryShareableVolumeVmInstanceRef(params *param.QueryParam) ([]view.ShareableVolumeVmInstanceRefInventoryView, error) {
	var resp []view.ShareableVolumeVmInstanceRefInventoryView
	return resp, cli.List("v1/volumes/vm-instances/refs", params, &resp)
}

// PageShareableVolumeVmInstanceRef Pagination
func (cli *ZSClient) PageShareableVolumeVmInstanceRef(params *param.QueryParam) ([]view.ShareableVolumeVmInstanceRefInventoryView, int, error) {
	var shareableVolumeVmInstanceRefs []view.ShareableVolumeVmInstanceRefInventoryView
	total, err := cli.Page("v1/volumes/vm-instances/refs", params, &shareableVolumeVmInstanceRefs)
	return shareableVolumeVmInstanceRefs, total, err
}
