// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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

func (cli *ZSClient) GetSlbOffering(uuid string) (*view.SlbOfferingInventoryView, error) {
	var resp view.SlbOfferingInventoryView
	if err := cli.Get("v1/instance-offerings/slb", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
