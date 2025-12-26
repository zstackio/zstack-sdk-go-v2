// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachPriceTableToAccount operates on PriceTableToAccount
func (cli *ZSClient) AttachPriceTableToAccount(params param.AttachPriceTableToAccountParam) (*view.AttachPriceTableToAccountEventView, error) {
	resp := view.AttachPriceTableToAccountEventView{}
	if err := cli.Post("v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
