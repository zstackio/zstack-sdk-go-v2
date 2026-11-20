// Copyright (c) ZStack.io, Inc.

package client

import (
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryHuaweiIMasterVpc queries HuaweiIMasterVpc list
func (cli *ZSClient) QueryHuaweiIMasterVpc(params *param.QueryParam) ([]view.HuaweiIMasterVpcInventoryView, error) {
	var resp []view.HuaweiIMasterVpcInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/vpcs", params, &resp)
}

func (cli *ZSClient) GetHuaweiIMasterVpc(uuid string) (*view.HuaweiIMasterVpcInventoryView, error) {
	var resp view.HuaweiIMasterVpcInventoryView
	if err := cli.Get("v1/sdn-controller/huawei-imaster/vpcs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHuaweiIMasterVpc Pagination
func (cli *ZSClient) PageHuaweiIMasterVpc(params *param.QueryParam) ([]view.HuaweiIMasterVpcInventoryView, int, error) {
	var huaweiIMasterVpcs []view.HuaweiIMasterVpcInventoryView
	total, err := cli.Page("v1/sdn-controller/huawei-imaster/vpcs", params, &huaweiIMasterVpcs)
	return huaweiIMasterVpcs, total, err
}
// DeleteHuaweiIMasterVpc sends the action with its default options.
func (cli *ZSClient) DeleteHuaweiIMasterVpc(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteHuaweiIMasterVpcWithParams(uuid, deleteMode, param.DeleteHuaweiIMasterVpcParamDetail{})
}

// DeleteHuaweiIMasterVpcWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteHuaweiIMasterVpcWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteHuaweiIMasterVpcParamDetail) error {
	query := url.Values{}
	if params.SdnControllerUuid != nil {
		query.Set("sdnControllerUuid", *params.SdnControllerUuid)
	}
	return cli.deleteResourceWithQuery("v1/sdn-controller/huawei-imaster/vpcs", uuid, deleteModeOrDefault(deleteMode, params.DeleteMode), query)
}
