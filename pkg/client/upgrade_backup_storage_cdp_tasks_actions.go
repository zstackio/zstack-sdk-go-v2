// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpgradeBackupStorageCdpTasks 操作UpgradeBackupStorageCdpTasks
func (cli *ZSClient) UpgradeBackupStorageCdpTasks(uuid string, params param.UpgradeBackupStorageCdpTasksParam) (*view.UpgradeBackupStorageCdpTasksEventView, error) {
	resp := view.UpgradeBackupStorageCdpTasksEventView{}
	if err := cli.Put("v1/cdp-task/upgrade/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

