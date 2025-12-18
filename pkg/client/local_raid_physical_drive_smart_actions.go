// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLocalRaidPhysicalDriveSmart 获取LocalRaidPhysicalDriveSmart详情
func (cli *ZSClient) GetLocalRaidPhysicalDriveSmart(uuid string) (*view.GetLocalRaidPhysicalDriveSmartView, error) {
	var resp view.GetLocalRaidPhysicalDriveSmartView
	if err := cli.Get("v1/storage-devices/local-raid/physical-drives/{uuid}/smart", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

