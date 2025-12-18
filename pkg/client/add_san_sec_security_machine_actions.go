// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSanSecSecurityMachine adds SanSecSecurityMachine
func (cli *ZSClient) AddSanSecSecurityMachine(params param.AddSanSecSecurityMachineParam) (*view.AddSecurityMachineEventView, error) {
	resp := view.AddSecurityMachineEventView{}
	if err := cli.Post("v1/security-machine/sanSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
