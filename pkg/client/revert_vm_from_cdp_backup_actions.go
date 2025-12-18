// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RevertVmFromCdpBackup 操作RevertVmFromCdpBackup
func (cli *ZSClient) RevertVmFromCdpBackup(uuid string, params param.RevertVmFromCdpBackupParam) (*view.RevertVmFromCdpBackupEventView, error) {
	resp := view.RevertVmFromCdpBackupEventView{}
	if err := cli.Put("v1/cdp-backups/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

