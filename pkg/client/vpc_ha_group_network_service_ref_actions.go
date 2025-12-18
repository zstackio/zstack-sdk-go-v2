// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcHaGroupNetworkServiceRef 查询VpcHaGroupNetworkServiceRef列表
func (cli *ZSClient) QueryVpcHaGroupNetworkServiceRef(params param.QueryParam) ([]view.QueryVpcHaGroupNetworkServiceRefView, error) {
	var resp []view.QueryVpcHaGroupNetworkServiceRefView
	return resp, cli.List("v1/vpc/hagroups/networkserviceref/", &params, &resp)
}

