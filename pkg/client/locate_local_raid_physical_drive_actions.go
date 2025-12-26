// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LocateLocalRaidPhysicalDrive operates on LocateLocalRaidPhysicalDrive
func (cli *ZSClient) LocateLocalRaidPhysicalDrive(uuid string, params param.LocateLocalRaidPhysicalDriveParam) (*view.LocateLocalRaidPhysicalDriveEventView, error) {
	resp := view.LocateLocalRaidPhysicalDriveEventView{}
	if err := cli.Put("v1/storage-devices/local-raid/physical-drives/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
