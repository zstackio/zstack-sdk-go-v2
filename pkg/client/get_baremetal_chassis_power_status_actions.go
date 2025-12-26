// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetBaremetalChassisPowerStatus gets BaremetalChassisPowerStatus by uuid
func (cli *ZSClient) GetBaremetalChassisPowerStatus(uuid string) (*view.GetBaremetalChassisPowerStatusView, error) {
	var resp view.GetBaremetalChassisPowerStatusView
	if err := cli.Get("v1/baremetal/chassis/{uuid}/powerstatus", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
