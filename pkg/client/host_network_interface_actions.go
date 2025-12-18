// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateHostNetworkInterface 更新HostNetworkInterface
func (cli *ZSClient) UpdateHostNetworkInterface(uuid string, params param.UpdateHostNetworkInterfaceParam) (*view.UpdateHostNetworkInterfaceEventView, error) {
	resp := view.UpdateHostNetworkInterfaceEventView{}
	if err := cli.Put("v1/hosts/nics/{interfaceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

