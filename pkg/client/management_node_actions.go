// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryManagementNode queries ManagementNode list
func (cli *ZSClient) QueryManagementNode(params *param.QueryParam) ([]view.ManagementNodeInventoryView, error) {
	var resp []view.ManagementNodeInventoryView
	return resp, cli.List("v1/management-nodes", params, &resp)
}

// PageManagementNode Pagination
func (cli *ZSClient) PageManagementNode(params *param.QueryParam) ([]view.ManagementNodeInventoryView, int, error) {
	var managementNodes []view.ManagementNodeInventoryView
	total, err := cli.Page("v1/management-nodes", params, &managementNodes)
	return managementNodes, total, err
}
