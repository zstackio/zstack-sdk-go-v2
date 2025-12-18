// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetL3NetworkRouterInterfaceIp 操作SetL3NetworkRouterInterfaceIp
func (cli *ZSClient) SetL3NetworkRouterInterfaceIp(params param.SetL3NetworkRouterInterfaceIpParam) (*view.SetL3NetworkRouterInterfaceIpEventView, error) {
	resp := view.SetL3NetworkRouterInterfaceIpEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/router-interface-ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

