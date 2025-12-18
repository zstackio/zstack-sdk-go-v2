// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBareMetal2Bonding 创建BareMetal2Bonding
func (cli *ZSClient) CreateBareMetal2Bonding(params param.CreateBareMetal2BondingParam) (*view.CreateBareMetal2BondingEventView, error) {
	resp := view.CreateBareMetal2BondingEventView{}
	if err := cli.Post("v1/baremetal2/chassis/bond", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

