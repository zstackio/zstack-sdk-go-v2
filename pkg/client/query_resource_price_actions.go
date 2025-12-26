// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryResourcePrice queries ResourcePrice list
func (cli *ZSClient) QueryResourcePrice(params *param.QueryParam) ([]view.PriceInventoryView, error) {
	var resp []view.PriceInventoryView
	return resp, cli.List("v1/billings/prices", params, &resp)
}
