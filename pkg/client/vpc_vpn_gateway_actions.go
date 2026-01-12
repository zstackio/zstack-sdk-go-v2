// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVpcVpnGateway updates VpcVpnGateway
func (cli *ZSClient) UpdateVpcVpnGateway(uuid string, params param.UpdateVpcVpnGatewayParam) (*view.VpcVpnGatewayInventoryView, error) {
	var resp view.UpdateVpcVpnGatewayEventView
	err := cli.PutWithSpec("v1/hybrid/vpc-vpn", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
