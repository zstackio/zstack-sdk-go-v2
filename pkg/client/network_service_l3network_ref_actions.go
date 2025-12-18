// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNetworkServiceL3NetworkRef 查询NetworkServiceL3NetworkRef列表
func (cli *ZSClient) QueryNetworkServiceL3NetworkRef(params param.QueryParam) ([]view.QueryNetworkServiceL3NetworkRefView, error) {
	var resp []view.QueryNetworkServiceL3NetworkRefView
	return resp, cli.List("v1/l3-networks/network-services/refs", &params, &resp)
}

