// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAccountResourceRef queries AccountResourceRef list
func (cli *ZSClient) QueryAccountResourceRef(params *param.QueryParam) ([]view.AccountResourceRefInventoryView, error) {
	var resp []view.AccountResourceRefInventoryView
	return resp, cli.List("v1/accounts/resources/refs", params, &resp)
}
