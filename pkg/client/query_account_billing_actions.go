// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAccountBilling queries AccountBilling list
func (cli *ZSClient) QueryAccountBilling(params *param.QueryParam) ([]view.BillingInventoryView, error) {
	var resp []view.BillingInventoryView
	return resp, cli.List("v1/billing/billings", params, &resp)
}
