// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncContainerManagementEndpoint 操作SyncContainerManagementEndpoint
func (cli *ZSClient) SyncContainerManagementEndpoint(uuid string, params param.SyncContainerManagementEndpointParam) (*view.SyncContainerManagementEndpointEventView, error) {
	resp := view.SyncContainerManagementEndpointEventView{}
	if err := cli.Put("v1/container/management/endpoint/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

