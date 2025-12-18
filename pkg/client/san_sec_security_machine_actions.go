// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSanSecSecurityMachine 更新SanSecSecurityMachine
func (cli *ZSClient) UpdateSanSecSecurityMachine(uuid string, params param.UpdateSanSecSecurityMachineParam) (*view.UpdateSecurityMachineEventView, error) {
	resp := view.UpdateSecurityMachineEventView{}
	if err := cli.Put("v1/security-machines/sanSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

