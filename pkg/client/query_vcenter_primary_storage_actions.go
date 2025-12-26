// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVCenterPrimaryStorage queries VCenterPrimaryStorage list
func (cli *ZSClient) QueryVCenterPrimaryStorage(params *param.QueryParam) ([]view.VCenterPrimaryStorageInventoryView, error) {
	var resp []view.VCenterPrimaryStorageInventoryView
	return resp, cli.List("v1/vcenters/primary-storage", params, &resp)
}
