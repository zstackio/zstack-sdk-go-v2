// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateFlkSecSecurityMachine 更新FlkSecSecurityMachine
func (cli *ZSClient) UpdateFlkSecSecurityMachine(uuid string, params param.UpdateFlkSecSecurityMachineParam) (*view.UpdateSecurityMachineEventView, error) {
	resp := view.UpdateSecurityMachineEventView{}
	if err := cli.Put("v1/security-machines/flkSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

