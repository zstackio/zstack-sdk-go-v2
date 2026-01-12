// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateNfvInst creates NfvInst
func (cli *ZSClient) CreateNfvInst(params param.CreateNfvInstParam) (*view.NfvInstInventoryView, error) {
	var resp view.CreateNfvInstEventView
	if err := cli.Post("v1/nfvinstgroup/inst", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryNfvInst queries NfvInst list
func (cli *ZSClient) QueryNfvInst(params *param.QueryParam) ([]view.NfvInstInventoryView, error) {
	var resp []view.NfvInstInventoryView
	return resp, cli.List("v1/vm-instances/appliances/nfvinst", params, &resp)
}

func (cli *ZSClient) GetNfvInst(uuid string) (*view.NfvInstInventoryView, error) {
	var resp view.NfvInstInventoryView
	if err := cli.Get("v1/vm-instances/appliances/nfvinst", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ReconnectNfvInst operates on NfvInst
func (cli *ZSClient) ReconnectNfvInst(vmInstanceUuid string, params param.ReconnectNfvInstParam) (*view.ApplianceVmInventoryView, error) {
	var resp view.ReconnectNfvInstEventView
	err := cli.PutWithSpec("v1/vm-instances/appliances/nfvinst", fmt.Sprintf(\"%s/actions\", vmInstanceUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
