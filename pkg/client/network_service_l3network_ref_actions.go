// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNetworkServiceL3NetworkRef queries NetworkServiceL3NetworkRef list
func (cli *ZSClient) QueryNetworkServiceL3NetworkRef(params *param.QueryParam) ([]view.NetworkServiceL3NetworkRefInventoryView, error) {
	var resp []view.NetworkServiceL3NetworkRefInventoryView
	return resp, cli.List("v1/l3-networks/network-services/refs", params, &resp)
}
