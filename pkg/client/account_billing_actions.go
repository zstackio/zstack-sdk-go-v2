// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAccountBilling 查询AccountBilling列表
func (cli *ZSClient) QueryAccountBilling(params param.QueryParam) ([]view.QueryAccountBillingView, error) {
	var resp []view.QueryAccountBillingView
	return resp, cli.List("v1/billing/billings", &params, &resp)
}

