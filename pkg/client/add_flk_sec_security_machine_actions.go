// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddFlkSecSecurityMachine 操作AddFlkSecSecurityMachine
func (cli *ZSClient) AddFlkSecSecurityMachine(params param.AddFlkSecSecurityMachineParam) (*view.AddSecurityMachineEventView, error) {
	resp := view.AddSecurityMachineEventView{}
	if err := cli.Post("v1/security-machine/flkSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

