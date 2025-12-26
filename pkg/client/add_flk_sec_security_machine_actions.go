// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddFlkSecSecurityMachine adds FlkSecSecurityMachine
func (cli *ZSClient) AddFlkSecSecurityMachine(params param.AddFlkSecSecurityMachineParam) (*view.AddSecurityMachineEventView, error) {
	resp := view.AddSecurityMachineEventView{}
	if err := cli.Post("v1/security-machine/flkSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
