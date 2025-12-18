// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVpcVRouterNetworkServiceState 操作SetVpcVRouterNetworkServiceState
func (cli *ZSClient) SetVpcVRouterNetworkServiceState(params param.SetVpcVRouterNetworkServiceStateParam) (*view.SetVpcVRouterNetworkServiceStateEventView, error) {
	resp := view.SetVpcVRouterNetworkServiceStateEventView{}
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/networkservicestate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

