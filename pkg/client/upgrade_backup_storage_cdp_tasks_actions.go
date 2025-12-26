// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpgradeBackupStorageCdpTasks operates on UpgradeBackupStorageCdpTasks
func (cli *ZSClient) UpgradeBackupStorageCdpTasks(uuid string, params param.UpgradeBackupStorageCdpTasksParam) (*view.UpgradeBackupStorageCdpTasksEventView, error) {
	resp := view.UpgradeBackupStorageCdpTasksEventView{}
	if err := cli.Put("v1/cdp-task/upgrade/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
