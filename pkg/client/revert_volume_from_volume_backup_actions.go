// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RevertVolumeFromVolumeBackup 操作RevertVolumeFromVolumeBackup
func (cli *ZSClient) RevertVolumeFromVolumeBackup(uuid string, params param.RevertVolumeFromVolumeBackupParam) (*view.RevertVolumeFromVolumeBackupEventView, error) {
	resp := view.RevertVolumeFromVolumeBackupEventView{}
	if err := cli.Put("v1/volume-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

