// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddContainerManagementEndpoint adds ContainerManagementEndpoint
func (cli *ZSClient) AddContainerManagementEndpoint(params param.AddContainerManagementEndpointParam) (*view.AddContainerManagementEndpointEventView, error) {
	resp := view.AddContainerManagementEndpointEventView{}
	if err := cli.Post("v1/container/management/endpoint", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
