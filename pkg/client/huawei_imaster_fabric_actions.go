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
	return cli.DeleteWithSpec("v1/sdn-controller/huawei-imaster/fabrics", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryHuaweiIMasterFabric queries HuaweiIMasterFabric list
func (cli *ZSClient) QueryHuaweiIMasterFabric(params *param.QueryParam) ([]view.HuaweiIMasterFabricInventoryView, error) {
	var resp []view.HuaweiIMasterFabricInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/fabrics", params, &resp)
}

func (cli *ZSClient) GetHuaweiIMasterFabric(uuid string) (*view.HuaweiIMasterFabricInventoryView, error) {
	var resp view.HuaweiIMasterFabricInventoryView
	if err := cli.Get("v1/sdn-controller/huawei-imaster/fabrics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
