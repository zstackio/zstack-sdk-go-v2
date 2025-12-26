// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSimulatorPrimaryStorage adds SimulatorPrimaryStorage
func (cli *ZSClient) AddSimulatorPrimaryStorage(params param.AddSimulatorPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
