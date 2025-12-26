// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncVpcVpnGatewayFromRemote operates on SyncVpcVpnGatewayFromRemote
func (cli *ZSClient) SyncVpcVpnGatewayFromRemote(uuid string, params param.SyncVpcVpnGatewayFromRemoteParam) (*view.SyncVpcVpnGatewayFromRemoteEventView, error) {
	resp := view.SyncVpcVpnGatewayFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/vpc-vpn/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
