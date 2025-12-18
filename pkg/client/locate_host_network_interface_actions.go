// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LocateHostNetworkInterface 操作LocateHostNetworkInterface
func (cli *ZSClient) LocateHostNetworkInterface(uuid string, params param.LocateHostNetworkInterfaceParam) (*view.LocateHostNetworkInterfaceEventView, error) {
	resp := view.LocateHostNetworkInterfaceEventView{}
	if err := cli.Put("v1/hosts/{hostUuid}/locate/network-interface", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

