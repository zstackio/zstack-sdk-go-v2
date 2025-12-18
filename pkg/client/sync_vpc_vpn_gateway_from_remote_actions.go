// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncVpcVpnGatewayFromRemote 操作SyncVpcVpnGatewayFromRemote
func (cli *ZSClient) SyncVpcVpnGatewayFromRemote(uuid string, params param.SyncVpcVpnGatewayFromRemoteParam) (*view.SyncVpcVpnGatewayFromRemoteEventView, error) {
	resp := view.SyncVpcVpnGatewayFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/vpc-vpn/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

