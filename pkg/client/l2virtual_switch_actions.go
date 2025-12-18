// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateL2VirtualSwitch 创建L2VirtualSwitch
func (cli *ZSClient) CreateL2VirtualSwitch(params param.CreateL2VirtualSwitchParam) (*view.CreateL2VirtualSwitchEventView, error) {
	resp := view.CreateL2VirtualSwitchEventView{}
	if err := cli.Post("v1/l2-networks/virtual-switch", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

