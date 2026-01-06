// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteHuaweiIMasterFabric deletes HuaweiIMasterFabric
func (cli *ZSClient) DeleteHuaweiIMasterFabric(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/fabrics/{uuid}", uuid, string(deleteMode))
}
// QueryHuaweiIMasterFabric queries HuaweiIMasterFabric list
func (cli *ZSClient) QueryHuaweiIMasterFabric(params *param.QueryParam) ([]view.HuaweiIMasterFabricInventoryView, error) {
	var resp []view.HuaweiIMasterFabricInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/fabrics", params, &resp)
}
