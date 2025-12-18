// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcVRouterNetworkServiceState gets VpcVRouterNetworkServiceState by uuid
func (cli *ZSClient) GetVpcVRouterNetworkServiceState(uuid string) (*view.GetVpcVRouterNetworkServiceStateView, error) {
	var resp view.GetVpcVRouterNetworkServiceStateView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/networkservicestate", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
