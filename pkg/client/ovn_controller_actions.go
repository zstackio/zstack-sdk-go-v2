// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryOvnController queries OvnController list
func (cli *ZSClient) QueryOvnController(params *param.QueryParam) ([]view.OvnControllerInventoryView, error) {
	var resp []view.OvnControllerInventoryView
	return resp, cli.List("v1/ovn-controllers", params, &resp)
}

// PageOvnController Pagination
func (cli *ZSClient) PageOvnController(params *param.QueryParam) ([]view.OvnControllerInventoryView, int, error) {
	var ovnControllers []view.OvnControllerInventoryView
	total, err := cli.Page("v1/ovn-controllers", params, &ovnControllers)
	return ovnControllers, total, err
}
