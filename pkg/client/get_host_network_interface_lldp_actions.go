// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostNetworkInterfaceLldp gets HostNetworkInterfaceLldp by uuid
func (cli *ZSClient) GetHostNetworkInterfaceLldp(uuid string) (*view.GetHostNetworkInterfaceLldpView, error) {
	var resp view.GetHostNetworkInterfaceLldpView
	if err := cli.Get("v1/hostNetworkInterface/lldp/{interfaceUuid}/info", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
