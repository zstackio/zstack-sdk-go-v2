// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryHuaweiIMasterVpc queries HuaweiIMasterVpc list
func (cli *ZSClient) QueryHuaweiIMasterVpc(params *param.QueryParam) ([]view.HuaweiIMasterVpcInventoryView, error) {
	var resp []view.HuaweiIMasterVpcInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/vpcs", params, &resp)
}
// DeleteHuaweiIMasterVpc deletes HuaweiIMasterVpc
func (cli *ZSClient) DeleteHuaweiIMasterVpc(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/vpcs/{uuid}", uuid, string(deleteMode))
}
