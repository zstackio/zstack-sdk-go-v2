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

// PageNfvInstOffering Pagination
func (cli *ZSClient) PageNfvInstOffering(params *param.QueryParam) ([]view.NfvInstOfferingInventoryView, int, error) {
	var nfvInstOfferings []view.NfvInstOfferingInventoryView
	total, err := cli.Page("v1/instance-offerings/nfvinst", params, &nfvInstOfferings)
	return nfvInstOfferings, total, err
}
// CreateNfvInstOffering creates NfvInstOffering
func (cli *ZSClient) CreateNfvInstOffering(params param.CreateNfvInstOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.Post("v1/instance-offerings/nfvinst", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
