// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVpnIkeConfig 创建VpnIkeConfig
func (cli *ZSClient) CreateVpnIkeConfig(params param.CreateVpnIkeConfigParam) (*view.CreateVpnIkeConfigEventView, error) {
	resp := view.CreateVpnIkeConfigEventView{}
	if err := cli.Post("v1/hybrid/vpn-connection/ike", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

