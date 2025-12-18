// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNativeCluster queries NativeCluster list
func (cli *ZSClient) QueryNativeCluster(params param.QueryParam) ([]view.NativeClusterInventoryView, error) {
	var resp []view.NativeClusterInventoryView
	return resp, cli.List("v1/container/native/cluster", &params, &resp)
}
