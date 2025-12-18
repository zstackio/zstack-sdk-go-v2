// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryHostNetworkBonding 查询HostNetworkBonding列表
func (cli *ZSClient) QueryHostNetworkBonding(params param.QueryParam) ([]view.QueryHostNetworkBondingView, error) {
	var resp []view.QueryHostNetworkBondingView
	return resp, cli.List("v1/hosts/bondings", &params, &resp)
}

