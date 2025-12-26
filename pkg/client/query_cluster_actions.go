// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryCluster queries Cluster list
func (cli *ZSClient) QueryCluster(params *param.QueryParam) ([]view.ClusterInventoryView, error) {
	var resp []view.ClusterInventoryView
	return resp, cli.List("v1/clusters", params, &resp)
}
