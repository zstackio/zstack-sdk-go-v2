// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryEmailMedia queries EmailMedia list
func (cli *ZSClient) QueryEmailMedia(params *param.QueryParam) ([]view.MediaInventoryView, error) {
	var resp []view.MediaInventoryView
	return resp, cli.List("v1/media/emails", params, &resp)
}
// CreateEmailMedia creates EmailMedia
func (cli *ZSClient) CreateEmailMedia(params param.CreateEmailMediaParam) (*view.MediaInventoryView, error) {
	var resp view.CreateMediaEventView
	if err := cli.Post("v1/media/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateEmailMedia updates EmailMedia
func (cli *ZSClient) UpdateEmailMedia(uuid string, params param.UpdateEmailMediaParam) (*view.EmailMediaInventoryView, error) {
	var resp view.UpdateEmailMediaEventView
	if err := cli.Put("v1/media/emails/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
