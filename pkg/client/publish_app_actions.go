// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPublishApp queries PublishApp list
func (cli *ZSClient) QueryPublishApp(params *param.QueryParam) ([]view.PublishAppInventoryView, error) {
	var resp []view.PublishAppInventoryView
	return resp, cli.List("v1/appcenter/app", params, &resp)
}

func (cli *ZSClient) GetPublishApp(uuid string) (*view.PublishAppInventoryView, error) {
	var resp view.PublishAppInventoryView
	if err := cli.Get("v1/appcenter/app", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
	err := cli.PutWithSpec("v1/appcenter/app", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePublishApp deletes PublishApp
func (cli *ZSClient) DeletePublishApp(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/appcenter/app", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
