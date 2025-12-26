// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVpcVRouterNetworkServiceState operates on SetVpcVRouterNetworkServiceState
func (cli *ZSClient) SetVpcVRouterNetworkServiceState(params param.SetVpcVRouterNetworkServiceStateParam) (*view.SetVpcVRouterNetworkServiceStateEventView, error) {
	resp := view.SetVpcVRouterNetworkServiceStateEventView{}
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/networkservicestate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
