// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNativeCluster queries NativeCluster list
func (cli *ZSClient) QueryNativeCluster(params *param.QueryParam) ([]view.NativeClusterInventoryView, error) {
	var resp []view.NativeClusterInventoryView
	return resp, cli.List("v1/container/native/cluster", params, &resp)
}

// PageNativeCluster Pagination
func (cli *ZSClient) PageNativeCluster(params *param.QueryParam) ([]view.NativeClusterInventoryView, int, error) {
	var nativeClusters []view.NativeClusterInventoryView
	total, err := cli.Page("v1/container/native/cluster", params, &nativeClusters)
	return nativeClusters, total, err
}
