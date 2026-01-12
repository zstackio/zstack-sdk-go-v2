// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectHost operates on Host
func (cli *ZSClient) ReconnectHost(uuid string, params param.ReconnectHostParam) (*view.HostInventoryView, error) {
	var resp view.ReconnectHostEventView
	err := cli.PutWithSpec("v1/hosts", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateHost updates Host
func (cli *ZSClient) UpdateHost(uuid string, params param.UpdateHostParam) (*view.HostInventoryView, error) {
	var resp view.UpdateHostEventView
	err := cli.PutWithSpec("v1/hosts", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteHost deletes Host
func (cli *ZSClient) DeleteHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/hosts", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryHost queries Host list
func (cli *ZSClient) QueryHost(params *param.QueryParam) ([]view.HostInventoryView, error) {
	var resp []view.HostInventoryView
	return resp, cli.List("v1/hosts", params, &resp)
}

func (cli *ZSClient) GetHost(uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.Get("v1/hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
