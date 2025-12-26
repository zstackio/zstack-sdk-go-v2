// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncContainerManagementEndpoint operates on SyncContainerManagementEndpoint
func (cli *ZSClient) SyncContainerManagementEndpoint(uuid string, params param.SyncContainerManagementEndpointParam) (*view.SyncContainerManagementEndpointEventView, error) {
	resp := view.SyncContainerManagementEndpointEventView{}
	if err := cli.Put("v1/container/management/endpoint/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
