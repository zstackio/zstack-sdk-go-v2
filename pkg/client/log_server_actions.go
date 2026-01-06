// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateLogServer updates LogServer
func (cli *ZSClient) UpdateLogServer(uuid string, params param.UpdateLogServerParam) (*view.LogServerInventoryView, error) {
	var resp view.UpdateLogServerEventView
	if err := cli.Put("v1/log/servers", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteLogServer deletes LogServer
func (cli *ZSClient) DeleteLogServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/log/servers", uuid, string(deleteMode))
}
// QueryLogServer queries LogServer list
func (cli *ZSClient) QueryLogServer(params *param.QueryParam) ([]view.LogServerInventoryView, error) {
	var resp []view.LogServerInventoryView
	return resp, cli.List("v1/log/servers", params, &resp)
}
// AddLogServer adds LogServer
func (cli *ZSClient) AddLogServer(params param.AddLogServerParam) (*view.LogServerInventoryView, error) {
	var resp view.AddLogServerEventView
	if err := cli.Post("v1/log/servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
