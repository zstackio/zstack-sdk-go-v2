// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryHuaweiIMasterVpc queries HuaweiIMasterVpc list
func (cli *ZSClient) QueryHuaweiIMasterVpc(ctx context.Context, params *param.QueryParam) ([]view.HuaweiIMasterVpcInventoryView, error) {
	var resp []view.HuaweiIMasterVpcInventoryView
	return resp, cli.List(ctx, "v1/sdn-controller/huawei-imaster/vpcs", params, &resp)
}

func (cli *ZSClient) GetHuaweiIMasterVpc(ctx context.Context, uuid string) (*view.HuaweiIMasterVpcInventoryView, error) {
	var resp view.HuaweiIMasterVpcInventoryView
	if err := cli.Get(ctx, "v1/sdn-controller/huawei-imaster/vpcs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHuaweiIMasterVpc Pagination
func (cli *ZSClient) PageHuaweiIMasterVpc(ctx context.Context, params *param.QueryParam) ([]view.HuaweiIMasterVpcInventoryView, int, error) {
	var huaweiIMasterVpcs []view.HuaweiIMasterVpcInventoryView
	total, err := cli.Page(ctx, "v1/sdn-controller/huawei-imaster/vpcs", params, &huaweiIMasterVpcs)
	return huaweiIMasterVpcs, total, err
}
// DeleteHuaweiIMasterVpc deletes HuaweiIMasterVpc
func (cli *ZSClient) DeleteHuaweiIMasterVpc(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/sdn-controller/huawei-imaster/vpcs", uuid, string(deleteMode))
}
