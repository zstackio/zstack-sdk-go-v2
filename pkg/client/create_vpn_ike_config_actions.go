// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVpnIkeConfig creates VpnIkeConfig
func (cli *ZSClient) CreateVpnIkeConfig(params param.CreateVpnIkeConfigParam) (*view.CreateVpnIkeConfigEventView, error) {
	resp := view.CreateVpnIkeConfigEventView{}
	if err := cli.Post("v1/hybrid/vpn-connection/ike", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
