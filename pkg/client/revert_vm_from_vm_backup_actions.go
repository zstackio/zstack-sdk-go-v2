// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RevertVmFromVmBackup operates on RevertVmFromVmBackup
func (cli *ZSClient) RevertVmFromVmBackup(uuid string, params param.RevertVmFromVmBackupParam) (*view.RevertVmFromVmBackupEventView, error) {
	resp := view.RevertVmFromVmBackupEventView{}
	if err := cli.Put("v1/vm-backups/{groupUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
