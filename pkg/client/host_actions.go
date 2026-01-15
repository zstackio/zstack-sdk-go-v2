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
	resp := view.HostInventoryView{}
	if err := cli.Put("v1/hosts", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateHost updates Host
func (cli *ZSClient) UpdateHost(uuid string, params param.UpdateHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.Put("v1/hosts", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteHost deletes Host
func (cli *ZSClient) DeleteHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts", uuid, string(deleteMode))
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

// PageHost Pagination
func (cli *ZSClient) PageHost(params *param.QueryParam) ([]view.HostInventoryView, int, error) {
	var hosts []view.HostInventoryView
	total, err := cli.Page("v1/hosts", params, &hosts)
	return hosts, total, err
}
