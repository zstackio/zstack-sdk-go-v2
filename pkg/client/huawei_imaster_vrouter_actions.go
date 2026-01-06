// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateHuaweiIMasterVRouter creates HuaweiIMasterVRouter
func (cli *ZSClient) CreateHuaweiIMasterVRouter(params param.CreateHuaweiIMasterVRouterParam) (*view.HuaweiIMasterVRouterInventoryView, error) {
	var resp view.CreateHuaweiIMasterVRouterEventView
	if err := cli.Post("v1/sdn-controller/huawei-imaster/vrouters", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryHuaweiIMasterVRouter queries HuaweiIMasterVRouter list
func (cli *ZSClient) QueryHuaweiIMasterVRouter(params *param.QueryParam) ([]view.HuaweiIMasterVRouterInventoryView, error) {
	var resp []view.HuaweiIMasterVRouterInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/vrouters", params, &resp)
}
// DeleteHuaweiIMasterVRouter deletes HuaweiIMasterVRouter
func (cli *ZSClient) DeleteHuaweiIMasterVRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/vrouters/{uuid}", uuid, string(deleteMode))
}
