// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSApplicationEndpoint 查询SNSApplicationEndpoint列表
func (cli *ZSClient) QuerySNSApplicationEndpoint(params param.QueryParam) ([]view.QuerySNSApplicationEndpointView, error) {
	var resp []view.QuerySNSApplicationEndpointView
	return resp, cli.List("v1/sns/application-endpoints", &params, &resp)
}

