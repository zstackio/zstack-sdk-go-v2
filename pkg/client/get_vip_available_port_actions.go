// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVipAvailablePort gets VipAvailablePort by uuid
func (cli *ZSClient) GetVipAvailablePort(uuid string) (*view.GetVipAvailablePortView, error) {
	var resp view.GetVipAvailablePortView
	if err := cli.Get("v1/vips/{vipUuid}/get-port-availability", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
