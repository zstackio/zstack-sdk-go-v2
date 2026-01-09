// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddV2VConversionHost adds V2VConversionHost
func (cli *ZSClient) AddV2VConversionHost(params param.AddV2VConversionHostParam) (*view.V2VConversionHostInventoryView, error) {
	var resp view.AddV2VConversionHostEventView
	if err := cli.Post("v1/v2v-conversion-hosts", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateV2VConversionHost updates V2VConversionHost
func (cli *ZSClient) UpdateV2VConversionHost(uuid string, params param.UpdateV2VConversionHostParam) (*view.V2VConversionHostInventoryView, error) {
	var resp view.UpdateV2VConversionHostEventView
	if err := cli.Put("v1/v2v-conversion-hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryV2VConversionHost queries V2VConversionHost list
func (cli *ZSClient) QueryV2VConversionHost(params *param.QueryParam) ([]view.V2VConversionHostInventoryView, error) {
	var resp []view.V2VConversionHostInventoryView
	return resp, cli.List("v1/v2v-conversion-hosts", params, &resp)
}

func (cli *ZSClient) GetV2VConversionHost(uuid string) (*view.V2VConversionHostInventoryView, error) {
	var resp view.V2VConversionHostInventoryView
	if err := cli.Get("v1/v2v-conversion-hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteV2VConversionHost deletes V2VConversionHost
func (cli *ZSClient) DeleteV2VConversionHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/v2v-conversion-hosts", uuid, string(deleteMode))
}
