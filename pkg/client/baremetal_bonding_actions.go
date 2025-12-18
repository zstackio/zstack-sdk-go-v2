// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBaremetalBonding 创建BaremetalBonding
func (cli *ZSClient) CreateBaremetalBonding(params param.CreateBaremetalBondingParam) (*view.CreateBaremetalBondingEventView, error) {
	resp := view.CreateBaremetalBondingEventView{}
	if err := cli.Post("v1/baremetal/network/bondings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

