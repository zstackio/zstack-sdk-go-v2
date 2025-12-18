// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExecuteGuestVmCommand 操作ExecuteGuestVmCommand
func (cli *ZSClient) ExecuteGuestVmCommand(params param.ExecuteGuestVmCommandParam) (*view.ExecuteGuestVmCommandEventView, error) {
	resp := view.ExecuteGuestVmCommandEventView{}
	if err := cli.Post("v1/vm-instances/commands/exec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

