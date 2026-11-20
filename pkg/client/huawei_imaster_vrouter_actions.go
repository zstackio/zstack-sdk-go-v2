// Copyright (c) ZStack.io, Inc.

package client

import (
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateHuaweiIMasterVRouter creates HuaweiIMasterVRouter
func (cli *ZSClient) CreateHuaweiIMasterVRouter(params param.CreateHuaweiIMasterVRouterParam) (*view.HuaweiIMasterVRouterInventoryView, error) {
	resp := view.HuaweiIMasterVRouterInventoryView{}
	if err := cli.Post("v1/sdn-controller/huawei-imaster/vrouters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageHuaweiIMasterVRouter Pagination
func (cli *ZSClient) PageHuaweiIMasterVRouter(params *param.QueryParam) ([]view.HuaweiIMasterVRouterInventoryView, int, error) {
	var huaweiIMasterVRouters []view.HuaweiIMasterVRouterInventoryView
	total, err := cli.Page("v1/sdn-controller/huawei-imaster/vrouters", params, &huaweiIMasterVRouters)
	return huaweiIMasterVRouters, total, err
}
// DeleteHuaweiIMasterVRouter sends the action with its default options.
func (cli *ZSClient) DeleteHuaweiIMasterVRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteHuaweiIMasterVRouterWithParams(uuid, deleteMode, param.DeleteHuaweiIMasterVRouterParamDetail{})
}

// DeleteHuaweiIMasterVRouterWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteHuaweiIMasterVRouterWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteHuaweiIMasterVRouterParamDetail) error {
	query := url.Values{}
	if params.SdnControllerUuid != nil {
		query.Set("sdnControllerUuid", *params.SdnControllerUuid)
	}
	return cli.deleteResourceWithQuery("v1/sdn-controller/huawei-imaster/vrouters", uuid, deleteModeOrDefault(deleteMode, params.DeleteMode), query)
}
