// Copyright (c) ZStack.io, Inc.

package client

import (
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
// DeleteHuaweiIMasterVpc deletes HuaweiIMasterVpc
func (cli *ZSClient) DeleteHuaweiIMasterVpc(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/vpcs", uuid, string(deleteMode))
}
