// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAddressPool queries AddressPool list
func (cli *ZSClient) QueryAddressPool(params *param.QueryParam) ([]view.AddressPoolInventoryView, error) {
	var resp []view.AddressPoolInventoryView
	return resp, cli.List("v1/l3-networks/address-pools", params, &resp)
}
