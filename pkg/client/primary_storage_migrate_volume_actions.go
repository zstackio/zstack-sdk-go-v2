// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PrimaryStorageMigrateVolume operates on PrimaryStorageMigrateVolume
func (cli *ZSClient) PrimaryStorageMigrateVolume(uuid string, params param.PrimaryStorageMigrateVolumeParam) (*view.PrimaryStorageMigrateVolumeEventView, error) {
	resp := view.PrimaryStorageMigrateVolumeEventView{}
	if err := cli.Put("v1/primary-storage/volumes/{volumeUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
