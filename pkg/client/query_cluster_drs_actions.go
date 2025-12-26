// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryClusterDRS queries ClusterDRS list
func (cli *ZSClient) QueryClusterDRS(params *param.QueryParam) ([]view.ClusterDRSInventoryView, error) {
	var resp []view.ClusterDRSInventoryView
	return resp, cli.List("v1/clusters/drs", params, &resp)
}
