// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryL2Network queries L2Network list
func (cli *ZSClient) QueryL2Network(params *param.QueryParam) ([]view.L2NetworkInventoryView, error) {
	var resp []view.L2NetworkInventoryView
	return resp, cli.List("v1/l2-networks", params, &resp)
}
