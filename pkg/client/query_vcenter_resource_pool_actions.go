// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVCenterResourcePool queries VCenterResourcePool list
func (cli *ZSClient) QueryVCenterResourcePool(params *param.QueryParam) ([]view.VCenterResourcePoolInventoryView, error) {
	var resp []view.VCenterResourcePoolInventoryView
	return resp, cli.List("v1/vcenters/clusters/resourcepools", params, &resp)
}
