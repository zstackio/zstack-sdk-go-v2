// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryCephPrimaryStoragePool queries CephPrimaryStoragePool list
func (cli *ZSClient) QueryCephPrimaryStoragePool(params *param.QueryParam) ([]view.CephPrimaryStoragePoolInventoryView, error) {
	var resp []view.CephPrimaryStoragePoolInventoryView
	return resp, cli.List("v1/primary-storage/ceph/pools", params, &resp)
}
