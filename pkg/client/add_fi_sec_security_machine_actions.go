// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddFiSecSecurityMachine 操作AddFiSecSecurityMachine
func (cli *ZSClient) AddFiSecSecurityMachine(params param.AddFiSecSecurityMachineParam) (*view.AddSecurityMachineEventView, error) {
	resp := view.AddSecurityMachineEventView{}
	if err := cli.Post("v1/security-machine/fiSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

