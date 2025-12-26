// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetInterdependentL3NetworksBackupStorages gets InterdependentL3NetworksBackupStorages by uuid
func (cli *ZSClient) GetInterdependentL3NetworksBackupStorages(uuid string) (*view.GetInterdependentL3NetworksBackupStoragesView, error) {
	var resp view.GetInterdependentL3NetworksBackupStoragesView
	if err := cli.Get("v1/backupStorage-l3networks/dependencies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
