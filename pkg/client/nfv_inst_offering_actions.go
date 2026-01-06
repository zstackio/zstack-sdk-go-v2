// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNfvInstOffering queries NfvInstOffering list
func (cli *ZSClient) QueryNfvInstOffering(params *param.QueryParam) ([]view.NfvInstOfferingInventoryView, error) {
	var resp []view.NfvInstOfferingInventoryView
	return resp, cli.List("v1/instance-offerings/nfvinst", params, &resp)
}
// CreateNfvInstOffering creates NfvInstOffering
func (cli *ZSClient) CreateNfvInstOffering(params param.CreateNfvInstOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	var resp view.CreateInstanceOfferingEventView
	if err := cli.Post("v1/instance-offerings/nfvinst", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
