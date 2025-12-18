// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAccountPriceTableRef 查询AccountPriceTableRef列表
func (cli *ZSClient) QueryAccountPriceTableRef(params param.QueryParam) ([]view.QueryAccountPriceTableRefView, error) {
	var resp []view.QueryAccountPriceTableRefView
	return resp, cli.List("v1/accounts/price-tables/refs", &params, &resp)
}

