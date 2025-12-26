// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateHostNetworkInterface updates HostNetworkInterface
func (cli *ZSClient) UpdateHostNetworkInterface(uuid string, params param.UpdateHostNetworkInterfaceParam) (*view.UpdateHostNetworkInterfaceEventView, error) {
	resp := view.UpdateHostNetworkInterfaceEventView{}
	if err := cli.Put("v1/hosts/nics/{interfaceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
