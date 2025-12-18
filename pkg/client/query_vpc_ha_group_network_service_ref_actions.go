// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcHaGroupNetworkServiceRef queries VpcHaGroupNetworkServiceRef list
func (cli *ZSClient) QueryVpcHaGroupNetworkServiceRef(params param.QueryParam) ([]view.VpcHaGroupNetworkServiceRefInventoryView, error) {
	var resp []view.VpcHaGroupNetworkServiceRefInventoryView
	return resp, cli.List("v1/vpc/hagroups/networkserviceref/", &params, &resp)
}
