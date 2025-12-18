// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateFiSecSecurityMachine 更新FiSecSecurityMachine
func (cli *ZSClient) UpdateFiSecSecurityMachine(uuid string, params param.UpdateFiSecSecurityMachineParam) (*view.UpdateSecurityMachineEventView, error) {
	resp := view.UpdateSecurityMachineEventView{}
	if err := cli.Put("v1/security-machines/fiSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

