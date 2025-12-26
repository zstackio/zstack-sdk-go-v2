// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSimulatorHost adds SimulatorHost
func (cli *ZSClient) AddSimulatorHost(params param.AddSimulatorHostParam) (*view.AddHostEventView, error) {
	resp := view.AddHostEventView{}
	if err := cli.Post("v1/hosts/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
