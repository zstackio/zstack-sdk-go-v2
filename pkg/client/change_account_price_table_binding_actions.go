// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeAccountPriceTableBinding changes AccountPriceTableBinding
func (cli *ZSClient) ChangeAccountPriceTableBinding(uuid string, params param.ChangeAccountPriceTableBindingParam) (*view.ChangeAccountPriceTableBindingEventView, error) {
	resp := view.ChangeAccountPriceTableBindingEventView{}
	if err := cli.Put("v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
