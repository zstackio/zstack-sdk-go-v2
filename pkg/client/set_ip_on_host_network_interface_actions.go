// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetIpOnHostNetworkInterface operates on SetIpOnHostNetworkInterface
func (cli *ZSClient) SetIpOnHostNetworkInterface(params param.SetIpOnHostNetworkInterfaceParam) (*view.SetIpOnHostNetworkInterfaceEventView, error) {
	resp := view.SetIpOnHostNetworkInterfaceEventView{}
	if err := cli.Post("v1/hosts/nics/{interfaceUuid}/ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
