// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryContainerManagementEndpoint 查询ContainerManagementEndpoint列表
func (cli *ZSClient) QueryContainerManagementEndpoint(params param.QueryParam) ([]view.QueryContainerManagementEndpointView, error) {
	var resp []view.QueryContainerManagementEndpointView
	return resp, cli.List("v1/container/management/endpoint", &params, &resp)
}

