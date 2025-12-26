// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// MergeDataOnBackupStorage operates on MergeDataOnBackupStorage
func (cli *ZSClient) MergeDataOnBackupStorage(uuid string, params param.MergeDataOnBackupStorageParam) (*view.MergeDataOnBackupStorageEventView, error) {
	resp := view.MergeDataOnBackupStorageEventView{}
	if err := cli.Put("v1/cdp-task/mergedata/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
