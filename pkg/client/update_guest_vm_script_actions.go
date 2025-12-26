// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateGuestVmScript updates GuestVmScript
func (cli *ZSClient) UpdateGuestVmScript(uuid string, params param.UpdateGuestVmScriptParam) (*view.UpdateGuestVmScriptEventView, error) {
	resp := view.UpdateGuestVmScriptEventView{}
	if err := cli.Put("v1/scripts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
