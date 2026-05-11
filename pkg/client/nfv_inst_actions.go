// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateNfvInst creates NfvInst
func (cli *ZSClient) CreateNfvInst(ctx context.Context, params param.CreateNfvInstParam) (*view.NfvInstInventoryView, error) {
	resp := view.NfvInstInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/nfvinstgroup/inst", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryNfvInst queries NfvInst list
func (cli *ZSClient) QueryNfvInst(ctx context.Context, params *param.QueryParam) ([]view.NfvInstInventoryView, error) {
	var resp []view.NfvInstInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/appliances/nfvinst", params, &resp)
}

func (cli *ZSClient) GetNfvInst(ctx context.Context, uuid string) (*view.NfvInstInventoryView, error) {
	var resp view.NfvInstInventoryView
	if err := cli.Get(ctx, "v1/vm-instances/appliances/nfvinst", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNfvInst Pagination
func (cli *ZSClient) PageNfvInst(ctx context.Context, params *param.QueryParam) ([]view.NfvInstInventoryView, int, error) {
	var nfvInsts []view.NfvInstInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/appliances/nfvinst", params, &nfvInsts)
	return nfvInsts, total, err
}
// ReconnectNfvInst operates on NfvInst
func (cli *ZSClient) ReconnectNfvInst(ctx context.Context, vmInstanceUuid string, params param.ReconnectNfvInstParam) (*view.ApplianceVmInventoryView, error) {
	resp := view.ApplianceVmInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/appliances/nfvinst", vmInstanceUuid, "", map[string]interface{}{
		"reconnectNfvInst": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
