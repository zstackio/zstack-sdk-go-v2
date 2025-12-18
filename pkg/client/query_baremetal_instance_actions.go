// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBaremetalInstance queries BaremetalInstance list
func (cli *ZSClient) QueryBaremetalInstance(params param.QueryParam) ([]view.BaremetalInstanceInventoryView, error) {
	var resp []view.BaremetalInstanceInventoryView
	return resp, cli.List("v1/baremetal/instances", &params, &resp)
}
