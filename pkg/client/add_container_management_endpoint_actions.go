// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddContainerManagementEndpoint 操作AddContainerManagementEndpoint
func (cli *ZSClient) AddContainerManagementEndpoint(params param.AddContainerManagementEndpointParam) (*view.AddContainerManagementEndpointEventView, error) {
	resp := view.AddContainerManagementEndpointEventView{}
	if err := cli.Post("v1/container/management/endpoint", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

