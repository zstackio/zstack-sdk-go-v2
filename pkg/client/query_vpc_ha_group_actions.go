// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcHaGroup queries VpcHaGroup list
func (cli *ZSClient) QueryVpcHaGroup(params *param.QueryParam) ([]view.VpcHaGroupInventoryView, error) {
	var resp []view.VpcHaGroupInventoryView
	return resp, cli.List("v1/vpc/hagroups", params, &resp)
}
