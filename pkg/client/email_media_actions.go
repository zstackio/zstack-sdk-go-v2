// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryEmailMedia queries EmailMedia list
func (cli *ZSClient) QueryEmailMedia(params *param.QueryParam) ([]view.MediaInventoryView, error) {
	var resp []view.MediaInventoryView
	return resp, cli.List("v1/media/emails", params, &resp)
}

func (cli *ZSClient) GetEmailMedia(uuid string) (*view.MediaInventoryView, error) {
	var resp view.MediaInventoryView
	if err := cli.Get("v1/media/emails", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
	err := cli.PutWithSpec("v1/media/emails", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
