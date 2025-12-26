// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVniRange queries VniRange list
func (cli *ZSClient) QueryVniRange(params *param.QueryParam) ([]view.VniRangeInventoryView, error) {
	var resp []view.VniRangeInventoryView
	return resp, cli.List("v1/l2-networks/vxlan-pool/vni-range", params, &resp)
}
