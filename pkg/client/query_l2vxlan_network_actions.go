// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryL2VxlanNetwork queries L2VxlanNetwork list
func (cli *ZSClient) QueryL2VxlanNetwork(params *param.QueryParam) ([]view.L2VxlanNetworkInventoryView, error) {
	var resp []view.L2VxlanNetworkInventoryView
	return resp, cli.List("v1/l2-networks/vxlan", params, &resp)
}
