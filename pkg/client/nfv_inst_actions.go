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
	resp := view.NfvInstInventoryView{}
	if err := cli.Post("v1/nfvinstgroup/inst", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageNfvInst Pagination
func (cli *ZSClient) PageNfvInst(params *param.QueryParam) ([]view.NfvInstInventoryView, int, error) {
	var nfvInsts []view.NfvInstInventoryView
	total, err := cli.Page("v1/vm-instances/appliances/nfvinst", params, &nfvInsts)
	return nfvInsts, total, err
}
// ReconnectNfvInst operates on NfvInst
func (cli *ZSClient) ReconnectNfvInst(vmInstanceUuid string, params param.ReconnectNfvInstParam) (*view.ApplianceVmInventoryView, error) {
	resp := view.ApplianceVmInventoryView{}
	if err := cli.Put("v1/vm-instances/appliances/nfvinst", vmInstanceUuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
