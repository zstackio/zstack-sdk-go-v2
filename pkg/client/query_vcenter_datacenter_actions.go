// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVCenterDatacenter queries VCenterDatacenter list
func (cli *ZSClient) QueryVCenterDatacenter(params param.QueryParam) ([]view.VCenterDatacenterInventoryView, error) {
	var resp []view.VCenterDatacenterInventoryView
	return resp, cli.List("v1/vcenters/datacenters", &params, &resp)
}
