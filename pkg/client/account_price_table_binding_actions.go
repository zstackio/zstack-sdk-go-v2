// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeAccountPriceTableBinding 操作AccountPriceTableBinding
func (cli *ZSClient) ChangeAccountPriceTableBinding(uuid string, params param.ChangeAccountPriceTableBindingParam) (*view.ChangeAccountPriceTableBindingEventView, error) {
	resp := view.ChangeAccountPriceTableBindingEventView{}
	if err := cli.Put("v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

