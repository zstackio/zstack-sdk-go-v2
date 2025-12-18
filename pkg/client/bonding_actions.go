// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBonding 创建Bonding
func (cli *ZSClient) CreateBonding(params param.CreateBondingParam) (*view.CreateBondingEventView, error) {
	resp := view.CreateBondingEventView{}
	if err := cli.Post("v1/hosts/bondings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

