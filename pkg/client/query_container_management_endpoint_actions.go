// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryContainerManagementEndpoint queries ContainerManagementEndpoint list
func (cli *ZSClient) QueryContainerManagementEndpoint(params param.QueryParam) ([]view.ContainerManagementEndpointInventoryView, error) {
	var resp []view.ContainerManagementEndpointInventoryView
	return resp, cli.List("v1/container/management/endpoint", &params, &resp)
}
