// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVCenterPrimaryStorage queries VCenterPrimaryStorage list
func (cli *ZSClient) QueryVCenterPrimaryStorage(params *param.QueryParam) ([]view.VCenterPrimaryStorageInventoryView, error) {
	var resp []view.VCenterPrimaryStorageInventoryView
	return resp, cli.List("v1/vcenters/primary-storage", params, &resp)
}

func (cli *ZSClient) GetVCenterPrimaryStorage(uuid string) (*view.VCenterPrimaryStorageInventoryView, error) {
	var resp view.VCenterPrimaryStorageInventoryView
	if err := cli.Get("v1/vcenters/primary-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
