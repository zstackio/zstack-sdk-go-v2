// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmBackup creates VmBackup
func (cli *ZSClient) CreateVmBackup(params param.CreateVmBackupParam) (*view.CreateVmBackupEventView, error) {
	resp := view.CreateVmBackupEventView{}
	if err := cli.Post("v1/volumes/{rootVolumeUuid}/vm-backups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
