// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryContainerManagementEndpoint queries ContainerManagementEndpoint list
func (cli *ZSClient) QueryContainerManagementEndpoint(params *param.QueryParam) ([]view.ContainerManagementEndpointInventoryView, error) {
	var resp []view.ContainerManagementEndpointInventoryView
	return resp, cli.List("v1/container/management/endpoint", params, &resp)
}
