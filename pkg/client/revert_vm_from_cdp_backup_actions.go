// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RevertVmFromCdpBackup operates on RevertVmFromCdpBackup
func (cli *ZSClient) RevertVmFromCdpBackup(uuid string, params param.RevertVmFromCdpBackupParam) (*view.RevertVmFromCdpBackupEventView, error) {
	resp := view.RevertVmFromCdpBackupEventView{}
	if err := cli.Put("v1/cdp-backups/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
