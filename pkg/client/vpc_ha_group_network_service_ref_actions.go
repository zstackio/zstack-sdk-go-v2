// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVpcHaGroupNetworkServiceRef queries VpcHaGroupNetworkServiceRef list
func (cli *ZSClient) QueryVpcHaGroupNetworkServiceRef(params *param.QueryParam) ([]view.VpcHaGroupNetworkServiceRefInventoryView, error) {
	var resp []view.VpcHaGroupNetworkServiceRefInventoryView
	return resp, cli.List("v1/vpc/hagroups/networkserviceref/", params, &resp)
}

// PageVpcHaGroupNetworkServiceRef Pagination
func (cli *ZSClient) PageVpcHaGroupNetworkServiceRef(params *param.QueryParam) ([]view.VpcHaGroupNetworkServiceRefInventoryView, int, error) {
	var vpcHaGroupNetworkServiceRefs []view.VpcHaGroupNetworkServiceRefInventoryView
	total, err := cli.Page("v1/vpc/hagroups/networkserviceref/", params, &vpcHaGroupNetworkServiceRefs)
	return vpcHaGroupNetworkServiceRefs, total, err
}
