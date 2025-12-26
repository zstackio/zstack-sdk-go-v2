// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSanSecSecurityMachine adds SanSecSecurityMachine
func (cli *ZSClient) AddSanSecSecurityMachine(params param.AddSanSecSecurityMachineParam) (*view.AddSecurityMachineEventView, error) {
	resp := view.AddSecurityMachineEventView{}
	if err := cli.Post("v1/security-machine/sanSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
