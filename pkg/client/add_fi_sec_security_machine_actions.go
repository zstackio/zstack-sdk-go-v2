// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddFiSecSecurityMachine adds FiSecSecurityMachine
func (cli *ZSClient) AddFiSecSecurityMachine(params param.AddFiSecSecurityMachineParam) (*view.AddSecurityMachineEventView, error) {
	resp := view.AddSecurityMachineEventView{}
	if err := cli.Post("v1/security-machine/fiSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
