// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryNativeCluster queries NativeCluster list
func (cli *ZSClient) QueryNativeCluster(params *param.QueryParam) ([]view.NativeClusterInventoryView, error) {
	var resp []view.NativeClusterInventoryView
	return resp, cli.List("v1/container/native/cluster", params, &resp)
}
