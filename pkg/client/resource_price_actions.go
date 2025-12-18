// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryResourcePrice 查询ResourcePrice列表
func (cli *ZSClient) QueryResourcePrice(params param.QueryParam) ([]view.QueryResourcePriceView, error) {
	var resp []view.QueryResourcePriceView
	return resp, cli.List("v1/billings/prices", &params, &resp)
}

