// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPortMirrorSession queries PortMirrorSession list
func (cli *ZSClient) QueryPortMirrorSession(params param.QueryParam) ([]view.PortMirrorSessionInventoryView, error) {
	var resp []view.PortMirrorSessionInventoryView
	return resp, cli.List("v1/port-mirrors/sessions", &params, &resp)
}
