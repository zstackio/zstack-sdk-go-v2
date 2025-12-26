// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddInfoSecSecurityMachine adds InfoSecSecurityMachine
func (cli *ZSClient) AddInfoSecSecurityMachine(params param.AddInfoSecSecurityMachineParam) (*view.AddSecurityMachineEventView, error) {
	resp := view.AddSecurityMachineEventView{}
	if err := cli.Post("v1/security-machine/infoSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
