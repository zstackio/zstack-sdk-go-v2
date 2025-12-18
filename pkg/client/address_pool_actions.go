// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAddressPool 查询AddressPool列表
func (cli *ZSClient) QueryAddressPool(params param.QueryParam) ([]view.QueryAddressPoolView, error) {
	var resp []view.QueryAddressPoolView
	return resp, cli.List("v1/l3-networks/address-pools", &params, &resp)
}

