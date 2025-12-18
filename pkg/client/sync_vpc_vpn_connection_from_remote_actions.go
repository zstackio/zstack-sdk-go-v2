// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncVpcVpnConnectionFromRemote 操作SyncVpcVpnConnectionFromRemote
func (cli *ZSClient) SyncVpcVpnConnectionFromRemote(uuid string, params param.SyncVpcVpnConnectionFromRemoteParam) (*view.SyncVpcVpnConnectionFromRemoteEventView, error) {
	resp := view.SyncVpcVpnConnectionFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/vpn-connection/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

