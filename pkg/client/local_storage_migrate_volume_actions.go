// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LocalStorageMigrateVolume 操作LocalStorageMigrateVolume
func (cli *ZSClient) LocalStorageMigrateVolume(uuid string, params param.LocalStorageMigrateVolumeParam) (*view.LocalStorageMigrateVolumeEventView, error) {
	resp := view.LocalStorageMigrateVolumeEventView{}
	if err := cli.Put("v1/primary-storage/local-storage/volumes/{volumeUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

