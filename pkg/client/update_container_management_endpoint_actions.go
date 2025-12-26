// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateContainerManagementEndpoint updates ContainerManagementEndpoint
func (cli *ZSClient) UpdateContainerManagementEndpoint(uuid string, params param.UpdateContainerManagementEndpointParam) (*view.UpdateContainerManagementEndpointEventView, error) {
	resp := view.UpdateContainerManagementEndpointEventView{}
	if err := cli.Put("v1/container/management/endpoint/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
