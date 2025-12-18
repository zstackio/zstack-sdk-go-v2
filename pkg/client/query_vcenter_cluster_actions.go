// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVCenterCluster queries VCenterCluster list
func (cli *ZSClient) QueryVCenterCluster(params param.QueryParam) ([]view.VCenterClusterInventoryView, error) {
	var resp []view.VCenterClusterInventoryView
	return resp, cli.List("v1/vcenters/clusters", &params, &resp)
}
