// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetServiceTypeOnHostNetworkInterface operates on SetServiceTypeOnHostNetworkInterface
func (cli *ZSClient) SetServiceTypeOnHostNetworkInterface(params param.SetServiceTypeOnHostNetworkInterfaceParam) (*view.SetServiceTypeOnHostNetworkInterfaceEventView, error) {
	resp := view.SetServiceTypeOnHostNetworkInterfaceEventView{}
	if err := cli.Post("v1/hosts/nics/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
