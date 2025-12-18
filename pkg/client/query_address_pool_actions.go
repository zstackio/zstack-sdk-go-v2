// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAddressPool queries AddressPool list
func (cli *ZSClient) QueryAddressPool(params param.QueryParam) ([]view.AddressPoolInventoryView, error) {
	var resp []view.AddressPoolInventoryView
	return resp, cli.List("v1/l3-networks/address-pools", &params, &resp)
}
