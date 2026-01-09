// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNfvInstOffering queries NfvInstOffering list
func (cli *ZSClient) QueryNfvInstOffering(params *param.QueryParam) ([]view.NfvInstOfferingInventoryView, error) {
	var resp []view.NfvInstOfferingInventoryView
	return resp, cli.List("v1/instance-offerings/nfvinst", params, &resp)
}

func (cli *ZSClient) GetNfvInstOffering(uuid string) (*view.NfvInstOfferingInventoryView, error) {
	var resp view.NfvInstOfferingInventoryView
	if err := cli.Get("v1/instance-offerings/nfvinst", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateNfvInstOffering creates NfvInstOffering
func (cli *ZSClient) CreateNfvInstOffering(params param.CreateNfvInstOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	var resp view.CreateInstanceOfferingEventView
	if err := cli.Post("v1/instance-offerings/nfvinst", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
