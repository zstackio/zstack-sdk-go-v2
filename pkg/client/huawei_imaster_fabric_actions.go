// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteHuaweiIMasterFabric deletes HuaweiIMasterFabric
func (cli *ZSClient) DeleteHuaweiIMasterFabric(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/fabrics", uuid, string(deleteMode))
}
// QueryHuaweiIMasterFabric queries HuaweiIMasterFabric list
func (cli *ZSClient) QueryHuaweiIMasterFabric(params *param.QueryParam) ([]view.HuaweiIMasterFabricInventoryView, error) {
	var resp []view.HuaweiIMasterFabricInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/fabrics", params, &resp)
}

// PageHuaweiIMasterFabric Pagination
func (cli *ZSClient) PageHuaweiIMasterFabric(params *param.QueryParam) ([]view.HuaweiIMasterFabricInventoryView, int, error) {
	var huaweiIMasterFabrics []view.HuaweiIMasterFabricInventoryView
	total, err := cli.Page("v1/sdn-controller/huawei-imaster/fabrics", params, &huaweiIMasterFabrics)
	return huaweiIMasterFabrics, total, err
}
