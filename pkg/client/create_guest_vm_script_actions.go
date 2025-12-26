// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateGuestVmScript creates GuestVmScript
func (cli *ZSClient) CreateGuestVmScript(params param.CreateGuestVmScriptParam) (*view.CreateGuestVmScriptEventView, error) {
	resp := view.CreateGuestVmScriptEventView{}
	if err := cli.Post("v1/scripts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
