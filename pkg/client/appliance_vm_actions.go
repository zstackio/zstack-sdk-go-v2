// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryApplianceVm queries ApplianceVm list
func (cli *ZSClient) QueryApplianceVm(params *param.QueryParam) ([]view.ApplianceVmInventoryView, error) {
	var resp []view.ApplianceVmInventoryView
	return resp, cli.List("v1/vm-instances/appliances", params, &resp)
}

func (cli *ZSClient) GetApplianceVm(uuid string) (*view.ApplianceVmInventoryView, error) {
	var resp view.ApplianceVmInventoryView
	if err := cli.Get("v1/vm-instances/appliances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
