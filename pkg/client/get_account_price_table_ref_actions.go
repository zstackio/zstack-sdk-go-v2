// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAccountPriceTableRef gets AccountPriceTableRef by uuid
func (cli *ZSClient) GetAccountPriceTableRef(uuid string) (*view.GetAccountPriceTableRefView, error) {
	var resp view.GetAccountPriceTableRefView
	if err := cli.Get("v1/billings/price-tables/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
