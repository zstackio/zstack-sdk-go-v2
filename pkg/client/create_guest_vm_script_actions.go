// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateGuestVmScript creates GuestVmScript
func (cli *ZSClient) CreateGuestVmScript(params param.CreateGuestVmScriptParam) (*view.CreateGuestVmScriptEventView, error) {
	resp := view.CreateGuestVmScriptEventView{}
	if err := cli.Post("v1/scripts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
