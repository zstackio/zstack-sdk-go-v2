// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmFromCdpBackup creates VmFromCdpBackup
func (cli *ZSClient) CreateVmFromCdpBackup(params param.CreateVmFromCdpBackupParam) (*view.CreateVmFromCdpBackupEventView, error) {
	resp := view.CreateVmFromCdpBackupEventView{}
	if err := cli.Post("v1/cdp-backups/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
