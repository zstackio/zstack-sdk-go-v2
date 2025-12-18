// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNetworkServiceProvider 查询NetworkServiceProvider列表
func (cli *ZSClient) QueryNetworkServiceProvider(params param.QueryParam) ([]view.QueryNetworkServiceProviderView, error) {
	var resp []view.QueryNetworkServiceProviderView
	return resp, cli.List("v1/network-services/providers", &params, &resp)
}

