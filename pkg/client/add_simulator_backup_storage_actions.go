// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSimulatorBackupStorage adds SimulatorBackupStorage
func (cli *ZSClient) AddSimulatorBackupStorage(params param.AddSimulatorBackupStorageParam) (*view.AddBackupStorageEventView, error) {
	resp := view.AddBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
