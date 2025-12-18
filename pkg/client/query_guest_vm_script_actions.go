// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGuestVmScript queries GuestVmScript list
func (cli *ZSClient) QueryGuestVmScript(params param.QueryParam) ([]view.GuestVmScriptInventoryView, error) {
	var resp []view.GuestVmScriptInventoryView
	return resp, cli.List("v1/scripts", &params, &resp)
}
