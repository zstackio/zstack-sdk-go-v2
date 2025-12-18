// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeHostNetworkInterfaceLldpMode changes HostNetworkInterfaceLldpMode
func (cli *ZSClient) ChangeHostNetworkInterfaceLldpMode(uuid string, params param.ChangeHostNetworkInterfaceLldpModeParam) (*view.ChangeHostNetworkInterfaceLldpModeEventView, error) {
	resp := view.ChangeHostNetworkInterfaceLldpModeEventView{}
	if err := cli.Put("v1/hostNetworkInterface/lldp/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
