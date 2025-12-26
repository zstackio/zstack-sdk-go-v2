// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncVpcVpnConnectionFromRemote operates on SyncVpcVpnConnectionFromRemote
func (cli *ZSClient) SyncVpcVpnConnectionFromRemote(uuid string, params param.SyncVpcVpnConnectionFromRemoteParam) (*view.SyncVpcVpnConnectionFromRemoteEventView, error) {
	resp := view.SyncVpcVpnConnectionFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/vpn-connection/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
