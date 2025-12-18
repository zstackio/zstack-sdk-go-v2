// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSecurityMachine 更新SecurityMachine
func (cli *ZSClient) UpdateSecurityMachine(uuid string, params param.UpdateSecurityMachineParam) (*view.UpdateSecurityMachineEventView, error) {
	resp := view.UpdateSecurityMachineEventView{}
	if err := cli.Put("v1/security-machines/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

