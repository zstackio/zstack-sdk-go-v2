// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExecuteGuestVmScript operates on ExecuteGuestVmScript
func (cli *ZSClient) ExecuteGuestVmScript(uuid string, params param.ExecuteGuestVmScriptParam) (*view.ExecuteGuestVmScriptEventView, error) {
	resp := view.ExecuteGuestVmScriptEventView{}
	if err := cli.Put("v1/scripts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
