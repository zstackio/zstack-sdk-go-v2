// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLocalRaidPhysicalDriveSmart gets LocalRaidPhysicalDriveSmart by uuid
func (cli *ZSClient) GetLocalRaidPhysicalDriveSmart(uuid string) (*view.GetLocalRaidPhysicalDriveSmartView, error) {
	var resp view.GetLocalRaidPhysicalDriveSmartView
	if err := cli.Get("v1/storage-devices/local-raid/physical-drives/{uuid}/smart", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
