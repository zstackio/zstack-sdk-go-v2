// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAlert queries Alert list
func (cli *ZSClient) QueryAlert(params *param.QueryParam) ([]view.AlertInventoryView, error) {
	var resp []view.AlertInventoryView
	return resp, cli.List("v1/monitoring/alerts", params, &resp)
}
// DeleteAlert deletes Alert
func (cli *ZSClient) DeleteAlert(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/alerts", uuid, string(deleteMode))
}
