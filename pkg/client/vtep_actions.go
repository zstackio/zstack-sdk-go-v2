// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVtep queries Vtep list
func (cli *ZSClient) QueryVtep(params *param.QueryParam) ([]view.VtepInventoryView, error) {
	var resp []view.VtepInventoryView
	return resp, cli.List("v1/l2-networks/vteps", params, &resp)
}

// PageVtep Pagination
func (cli *ZSClient) PageVtep(params *param.QueryParam) ([]view.VtepInventoryView, int, error) {
	var vteps []view.VtepInventoryView
	total, err := cli.Page("v1/l2-networks/vteps", params, &vteps)
	return vteps, total, err
}
