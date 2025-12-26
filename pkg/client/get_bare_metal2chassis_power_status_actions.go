// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetBareMetal2ChassisPowerStatus gets BareMetal2ChassisPowerStatus by uuid
func (cli *ZSClient) GetBareMetal2ChassisPowerStatus(uuid string) (*view.GetBareMetal2ChassisPowerStatusView, error) {
	var resp view.GetBareMetal2ChassisPowerStatusView
	if err := cli.Get("v1/baremetal2/chassis/{uuid}/powerstatus", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
