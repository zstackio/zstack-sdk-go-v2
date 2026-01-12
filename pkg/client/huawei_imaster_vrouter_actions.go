// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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

func (cli *ZSClient) GetHuaweiIMasterVRouter(uuid string) (*view.HuaweiIMasterVRouterInventoryView, error) {
	var resp view.HuaweiIMasterVRouterInventoryView
	if err := cli.Get("v1/sdn-controller/huawei-imaster/vrouters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteHuaweiIMasterVRouter deletes HuaweiIMasterVRouter
func (cli *ZSClient) DeleteHuaweiIMasterVRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/sdn-controller/huawei-imaster/vrouters", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
