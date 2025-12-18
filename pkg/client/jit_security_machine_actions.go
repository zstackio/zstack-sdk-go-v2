// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateJitSecurityMachine 更新JitSecurityMachine
func (cli *ZSClient) UpdateJitSecurityMachine(uuid string, params param.UpdateJitSecurityMachineParam) (*view.UpdateSecurityMachineEventView, error) {
	resp := view.UpdateSecurityMachineEventView{}
	if err := cli.Put("v1/security-machines/jida/auth-gateway/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

