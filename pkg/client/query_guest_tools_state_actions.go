// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGuestToolsState queries GuestToolsState list
func (cli *ZSClient) QueryGuestToolsState(params param.QueryParam) ([]view.GuestToolsStateInventoryView, error) {
	var resp []view.GuestToolsStateInventoryView
	return resp, cli.List("v1/guesttools", &params, &resp)
}
