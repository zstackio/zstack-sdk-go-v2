// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPublishApp queries PublishApp list
func (cli *ZSClient) QueryPublishApp(params *param.QueryParam) ([]view.PublishAppInventoryView, error) {
	var resp []view.PublishAppInventoryView
	return resp, cli.List("v1/appcenter/app", params, &resp)
}
// PublishApp operates on PublishApp
func (cli *ZSClient) PublishApp(params param.PublishAppParam) (*view.PublishAppInventoryView, error) {
	var resp view.PublishAppEventView
	if err := cli.Post("v1/appcenter/app", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdatePublishApp updates PublishApp
func (cli *ZSClient) UpdatePublishApp(uuid string, params param.UpdatePublishAppParam) (*view.PublishAppInventoryView, error) {
	var resp view.UpdatePublishAppEventView
	if err := cli.Put("v1/appcenter/app/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePublishApp deletes PublishApp
func (cli *ZSClient) DeletePublishApp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/app/{uuid}", uuid, string(deleteMode))
}
