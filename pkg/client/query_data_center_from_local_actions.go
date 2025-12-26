// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryDataCenterFromLocal queries DataCenterFromLocal list
func (cli *ZSClient) QueryDataCenterFromLocal(params *param.QueryParam) ([]view.DataCenterInventoryView, error) {
	var resp []view.DataCenterInventoryView
	return resp, cli.List("v1/hybrid/data-center", params, &resp)
}
