// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySlbGroup queries SlbGroup list
func (cli *ZSClient) QuerySlbGroup(params *param.QueryParam) ([]view.SlbGroupInventoryView, error) {
	var resp []view.SlbGroupInventoryView
	return resp, cli.List("v1/load-balancers/slb/groups", params, &resp)
}
