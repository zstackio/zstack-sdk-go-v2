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

// PageVCenterPrimaryStorage Pagination
func (cli *ZSClient) PageVCenterPrimaryStorage(params *param.QueryParam) ([]view.VCenterPrimaryStorageInventoryView, int, error) {
	var vCenterPrimaryStorages []view.VCenterPrimaryStorageInventoryView
	total, err := cli.Page("v1/vcenters/primary-storage", params, &vCenterPrimaryStorages)
	return vCenterPrimaryStorages, total, err
}
