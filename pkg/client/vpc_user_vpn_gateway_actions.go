// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVpcUserVpnGateway updates VpcUserVpnGateway
func (cli *ZSClient) UpdateVpcUserVpnGateway(uuid string, params param.UpdateVpcUserVpnGatewayParam) (*view.VpcUserVpnGatewayInventoryView, error) {
	var resp view.UpdateVpcUserVpnGatewayEventView
	err := cli.PutWithSpec("v1/hybrid/user-vpn", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
