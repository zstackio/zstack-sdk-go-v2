// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmFromVolumeBackup creates VmFromVolumeBackup
func (cli *ZSClient) CreateVmFromVolumeBackup(params param.CreateVmFromVolumeBackupParam) (*view.CreateVmFromVolumeBackupEventView, error) {
	resp := view.CreateVmFromVolumeBackupEventView{}
	if err := cli.Post("v1/vm-instances/from/vm-backup/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
