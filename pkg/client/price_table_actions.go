// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdatePriceTable 更新PriceTable
func (cli *ZSClient) UpdatePriceTable(uuid string, params param.UpdatePriceTableParam) (*view.UpdatePriceTableEventView, error) {
	resp := view.UpdatePriceTableEventView{}
	if err := cli.Put("v1/billings/price-tables/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

