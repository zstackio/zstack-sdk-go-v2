// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddJitSecurityMachine adds JitSecurityMachine
func (cli *ZSClient) AddJitSecurityMachine(params param.AddJitSecurityMachineParam) (*view.AddSecurityMachineEventView, error) {
	resp := view.AddSecurityMachineEventView{}
	if err := cli.Post("v1/security-machine/jida/auth-gateway", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
