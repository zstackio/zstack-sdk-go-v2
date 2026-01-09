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

func (cli *ZSClient) GetShareableVolumeVmInstanceRef(uuid string) (*view.ShareableVolumeVmInstanceRefInventoryView, error) {
	var resp view.ShareableVolumeVmInstanceRefInventoryView
	if err := cli.Get("v1/volumes/vm-instances/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
