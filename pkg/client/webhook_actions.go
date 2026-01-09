// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryWebhook queries Webhook list
func (cli *ZSClient) QueryWebhook(params *param.QueryParam) ([]view.WebhookInventoryView, error) {
	var resp []view.WebhookInventoryView
	return resp, cli.List("v1/web-hooks", params, &resp)
}

func (cli *ZSClient) GetWebhook(uuid string) (*view.WebhookInventoryView, error) {
	var resp view.WebhookInventoryView
	if err := cli.Get("v1/web-hooks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateWebhook updates Webhook
func (cli *ZSClient) UpdateWebhook(uuid string, params param.UpdateWebhookParam) (*view.WebhookInventoryView, error) {
	var resp view.UpdateWebhookEventView
	if err := cli.Put("v1/web-hooks/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteWebhook deletes Webhook
func (cli *ZSClient) DeleteWebhook(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/web-hooks", uuid, string(deleteMode))
}
// CreateWebhook creates Webhook
func (cli *ZSClient) CreateWebhook(params param.CreateWebhookParam) (*view.WebhookInventoryView, error) {
	var resp view.CreateWebhookEventView
	if err := cli.Post("v1/web-hooks", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
