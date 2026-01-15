// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePortMirrorSession creates PortMirrorSession
func (cli *ZSClient) CreatePortMirrorSession(params param.CreatePortMirrorSessionParam) (*view.PortMirrorSessionInventoryView, error) {
	resp := view.PortMirrorSessionInventoryView{}
	if err := cli.Post("v1/port-mirrors/sessions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePortMirrorSession deletes PortMirrorSession
func (cli *ZSClient) DeletePortMirrorSession(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-mirrors/sessons", uuid, string(deleteMode))
}
// QueryPortMirrorSession queries PortMirrorSession list
func (cli *ZSClient) QueryPortMirrorSession(params *param.QueryParam) ([]view.PortMirrorSessionInventoryView, error) {
	var resp []view.PortMirrorSessionInventoryView
	return resp, cli.List("v1/port-mirrors/sessions", params, &resp)
}

// PagePortMirrorSession Pagination
func (cli *ZSClient) PagePortMirrorSession(params *param.QueryParam) ([]view.PortMirrorSessionInventoryView, int, error) {
	var portMirrorSessions []view.PortMirrorSessionInventoryView
	total, err := cli.Page("v1/port-mirrors/sessions", params, &portMirrorSessions)
	return portMirrorSessions, total, err
}
