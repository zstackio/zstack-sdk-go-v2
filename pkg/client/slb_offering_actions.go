// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSlbOffering creates SlbOffering
func (cli *ZSClient) CreateSlbOffering(params param.CreateSlbOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	var resp view.CreateInstanceOfferingEventView
	if err := cli.Post("v1/instance-offerings/slb", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySlbOffering queries SlbOffering list
func (cli *ZSClient) QuerySlbOffering(params *param.QueryParam) ([]view.SlbOfferingInventoryView, error) {
	var resp []view.SlbOfferingInventoryView
	return resp, cli.List("v1/instance-offerings/slb", params, &resp)
}
