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

// PageEmailMedia Pagination
func (cli *ZSClient) PageEmailMedia(params *param.QueryParam) ([]view.MediaInventoryView, int, error) {
	var emailMedias []view.MediaInventoryView
	total, err := cli.Page("v1/media/emails", params, &emailMedias)
	return emailMedias, total, err
}
// CreateEmailMedia creates EmailMedia
func (cli *ZSClient) CreateEmailMedia(params param.CreateEmailMediaParam) (*view.MediaInventoryView, error) {
	resp := view.MediaInventoryView{}
	if err := cli.Post("v1/media/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateEmailMedia updates EmailMedia
func (cli *ZSClient) UpdateEmailMedia(uuid string, params param.UpdateEmailMediaParam) (*view.EmailMediaInventoryView, error) {
	resp := view.EmailMediaInventoryView{}
	if err := cli.Put("v1/media/emails", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
