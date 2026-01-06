// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectZdfs operates on Zdfs
func (cli *ZSClient) ReconnectZdfs(uuid string, params param.ReconnectZdfsParam) (*view.ZdfsInventoryView, error) {
	var resp view.ReconnectZdfsEventView
	if err := cli.Put("v1/zdfs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryZdfs queries Zdfs list
func (cli *ZSClient) QueryZdfs(params *param.QueryParam) ([]view.ZdfsInventoryView, error) {
	var resp []view.ZdfsInventoryView
	return resp, cli.List("v1/zdfs", params, &resp)
}
