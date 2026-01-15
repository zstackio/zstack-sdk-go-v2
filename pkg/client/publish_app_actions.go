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

// PagePublishApp Pagination
func (cli *ZSClient) PagePublishApp(params *param.QueryParam) ([]view.PublishAppInventoryView, int, error) {
	var publishApps []view.PublishAppInventoryView
	total, err := cli.Page("v1/appcenter/app", params, &publishApps)
	return publishApps, total, err
}
// PublishApp operates on PublishApp
func (cli *ZSClient) PublishApp(params param.PublishAppParam) (*view.PublishAppInventoryView, error) {
	resp := view.PublishAppInventoryView{}
	if err := cli.Post("v1/appcenter/app", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePublishApp updates PublishApp
func (cli *ZSClient) UpdatePublishApp(uuid string, params param.UpdatePublishAppParam) (*view.PublishAppInventoryView, error) {
	resp := view.PublishAppInventoryView{}
	if err := cli.Put("v1/appcenter/app", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePublishApp deletes PublishApp
func (cli *ZSClient) DeletePublishApp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/app", uuid, string(deleteMode))
}
