// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncVpcUserVpnGatewayFromRemote operates on SyncVpcUserVpnGatewayFromRemote
func (cli *ZSClient) SyncVpcUserVpnGatewayFromRemote(uuid string, params param.SyncVpcUserVpnGatewayFromRemoteParam) (*view.SyncVpcUserVpnGatewayFromRemoteEventView, error) {
	resp := view.SyncVpcUserVpnGatewayFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/user-vpn/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
