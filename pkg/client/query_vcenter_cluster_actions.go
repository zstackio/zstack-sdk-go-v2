// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVCenterCluster queries VCenterCluster list
func (cli *ZSClient) QueryVCenterCluster(params *param.QueryParam) ([]view.VCenterClusterInventoryView, error) {
	var resp []view.VCenterClusterInventoryView
	return resp, cli.List("v1/vcenters/clusters", params, &resp)
}
