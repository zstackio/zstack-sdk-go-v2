// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmFromVmBackup creates VmFromVmBackup
func (cli *ZSClient) CreateVmFromVmBackup(params param.CreateVmFromVmBackupParam) (*view.CreateVmFromVmBackupEventView, error) {
	resp := view.CreateVmFromVmBackupEventView{}
	if err := cli.Post("v1/vm-instances/from/vm-backups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
