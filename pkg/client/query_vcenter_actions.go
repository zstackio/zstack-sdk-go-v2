// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVCenter queries VCenter list
func (cli *ZSClient) QueryVCenter(params param.QueryParam) ([]view.VCenterInventoryView, error) {
	var resp []view.VCenterInventoryView
	return resp, cli.List("v1/vcenters", &params, &resp)
}
