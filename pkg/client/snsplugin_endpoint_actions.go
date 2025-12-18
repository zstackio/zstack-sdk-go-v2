// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSPluginEndpoint 查询SNSPluginEndpoint列表
func (cli *ZSClient) QuerySNSPluginEndpoint(params param.QueryParam) ([]view.QuerySNSPluginEndpointView, error) {
	var resp []view.QuerySNSPluginEndpointView
	return resp, cli.List("v1/sns/application-endpoints/plugin", &params, &resp)
}

