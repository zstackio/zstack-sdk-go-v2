// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryConnectionAccessPointFromLocal queries ConnectionAccessPointFromLocal list
func (cli *ZSClient) QueryConnectionAccessPointFromLocal(params *param.QueryParam) ([]view.ConnectionAccessPointInventoryView, error) {
	var resp []view.ConnectionAccessPointInventoryView
	return resp, cli.List("v1/hybrid/aliyun/access-point", params, &resp)
}
