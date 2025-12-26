// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LocalStorageMigrateVolume operates on LocalStorageMigrateVolume
func (cli *ZSClient) LocalStorageMigrateVolume(uuid string, params param.LocalStorageMigrateVolumeParam) (*view.LocalStorageMigrateVolumeEventView, error) {
	resp := view.LocalStorageMigrateVolumeEventView{}
	if err := cli.Put("v1/primary-storage/local-storage/volumes/{volumeUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
