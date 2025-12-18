// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetL3NetworkRouterInterfaceIp gets L3NetworkRouterInterfaceIp by uuid
func (cli *ZSClient) GetL3NetworkRouterInterfaceIp(uuid string) (*view.GetL3NetworkRouterInterfaceIpView, error) {
	var resp view.GetL3NetworkRouterInterfaceIpView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/router-interface-ip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
