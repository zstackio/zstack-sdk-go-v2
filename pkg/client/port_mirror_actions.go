// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePortMirror creates PortMirror
func (cli *ZSClient) CreatePortMirror(params param.CreatePortMirrorParam) (*view.PortMirrorInventoryView, error) {
	var resp view.CreatePortMirrorEventView
	if err := cli.Post("v1/port-mirrors", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryPortMirror queries PortMirror list
func (cli *ZSClient) QueryPortMirror(params *param.QueryParam) ([]view.PortMirrorInventoryView, error) {
	var resp []view.PortMirrorInventoryView
	return resp, cli.List("v1/port-mirrors", params, &resp)
}

func (cli *ZSClient) GetPortMirror(uuid string) (*view.PortMirrorInventoryView, error) {
	var resp view.PortMirrorInventoryView
	if err := cli.Get("v1/port-mirrors", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePortMirror deletes PortMirror
func (cli *ZSClient) DeletePortMirror(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/port-mirrors", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdatePortMirror updates PortMirror
func (cli *ZSClient) UpdatePortMirror(uuid string, params param.UpdatePortMirrorParam) (*view.PortMirrorInventoryView, error) {
	var resp view.UpdatePortMirrorEventView
	err := cli.PutWithSpec("v1/port-mirrors", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
