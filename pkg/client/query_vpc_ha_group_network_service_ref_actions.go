// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcHaGroupNetworkServiceRef queries VpcHaGroupNetworkServiceRef list
func (cli *ZSClient) QueryVpcHaGroupNetworkServiceRef(params *param.QueryParam) ([]view.VpcHaGroupNetworkServiceRefInventoryView, error) {
	var resp []view.VpcHaGroupNetworkServiceRefInventoryView
	return resp, cli.List("v1/vpc/hagroups/networkserviceref/", params, &resp)
}
