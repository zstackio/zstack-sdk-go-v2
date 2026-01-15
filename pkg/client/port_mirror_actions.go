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
	resp := view.PortMirrorInventoryView{}
	if err := cli.Post("v1/port-mirrors", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPortMirror queries PortMirror list
func (cli *ZSClient) QueryPortMirror(params *param.QueryParam) ([]view.PortMirrorInventoryView, error) {
	var resp []view.PortMirrorInventoryView
	return resp, cli.List("v1/port-mirrors", params, &resp)
}

// PagePortMirror Pagination
func (cli *ZSClient) PagePortMirror(params *param.QueryParam) ([]view.PortMirrorInventoryView, int, error) {
	var portMirrors []view.PortMirrorInventoryView
	total, err := cli.Page("v1/port-mirrors", params, &portMirrors)
	return portMirrors, total, err
}
// DeletePortMirror deletes PortMirror
func (cli *ZSClient) DeletePortMirror(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-mirrors", uuid, string(deleteMode))
}
// UpdatePortMirror updates PortMirror
func (cli *ZSClient) UpdatePortMirror(uuid string, params param.UpdatePortMirrorParam) (*view.PortMirrorInventoryView, error) {
	resp := view.PortMirrorInventoryView{}
	if err := cli.Put("v1/port-mirrors", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
