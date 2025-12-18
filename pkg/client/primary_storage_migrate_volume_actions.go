// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PrimaryStorageMigrateVolume 操作PrimaryStorageMigrateVolume
func (cli *ZSClient) PrimaryStorageMigrateVolume(uuid string, params param.PrimaryStorageMigrateVolumeParam) (*view.PrimaryStorageMigrateVolumeEventView, error) {
	resp := view.PrimaryStorageMigrateVolumeEventView{}
	if err := cli.Put("v1/primary-storage/volumes/{volumeUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

