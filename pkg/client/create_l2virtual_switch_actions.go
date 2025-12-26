// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateL2VirtualSwitch creates L2VirtualSwitch
func (cli *ZSClient) CreateL2VirtualSwitch(params param.CreateL2VirtualSwitchParam) (*view.CreateL2VirtualSwitchEventView, error) {
	resp := view.CreateL2VirtualSwitchEventView{}
	if err := cli.Post("v1/l2-networks/virtual-switch", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
