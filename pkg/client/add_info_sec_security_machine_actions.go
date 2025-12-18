// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddInfoSecSecurityMachine 操作AddInfoSecSecurityMachine
func (cli *ZSClient) AddInfoSecSecurityMachine(params param.AddInfoSecSecurityMachineParam) (*view.AddSecurityMachineEventView, error) {
	resp := view.AddSecurityMachineEventView{}
	if err := cli.Post("v1/security-machine/infoSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

