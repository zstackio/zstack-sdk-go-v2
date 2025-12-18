// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateGuestVmScript updates GuestVmScript
func (cli *ZSClient) UpdateGuestVmScript(uuid string, params param.UpdateGuestVmScriptParam) (*view.UpdateGuestVmScriptEventView, error) {
	resp := view.UpdateGuestVmScriptEventView{}
	if err := cli.Put("v1/scripts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
