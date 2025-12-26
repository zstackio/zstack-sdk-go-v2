// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LocateHostNetworkInterface operates on LocateHostNetworkInterface
func (cli *ZSClient) LocateHostNetworkInterface(uuid string, params param.LocateHostNetworkInterfaceParam) (*view.LocateHostNetworkInterfaceEventView, error) {
	resp := view.LocateHostNetworkInterfaceEventView{}
	if err := cli.Put("v1/hosts/{hostUuid}/locate/network-interface", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
