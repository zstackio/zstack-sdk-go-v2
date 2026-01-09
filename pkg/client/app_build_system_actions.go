// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAppBuildSystem adds AppBuildSystem
func (cli *ZSClient) AddAppBuildSystem(params param.AddAppBuildSystemParam) (*view.AppBuildSystemInventoryView, error) {
	var resp view.AddAppBuildSystemEventView
	if err := cli.Post("v1/appcenter/buildsystem", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAppBuildSystem queries AppBuildSystem list
func (cli *ZSClient) QueryAppBuildSystem(params *param.QueryParam) ([]view.AppBuildSystemInventoryView, error) {
	var resp []view.AppBuildSystemInventoryView
	return resp, cli.List("v1/appcenter/buildsystem", params, &resp)
}

func (cli *ZSClient) GetAppBuildSystem(uuid string) (*view.AppBuildSystemInventoryView, error) {
	var resp view.AppBuildSystemInventoryView
	if err := cli.Get("v1/appcenter/buildsystem", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ReconnectAppBuildSystem operates on AppBuildSystem
func (cli *ZSClient) ReconnectAppBuildSystem(uuid string, params param.ReconnectAppBuildSystemParam) (*view.AppBuildSystemInventoryView, error) {
	var resp view.ReconnectAppBuildSystemEventView
	if err := cli.Put("v1/appcenter/buildsystem/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateAppBuildSystem updates AppBuildSystem
func (cli *ZSClient) UpdateAppBuildSystem(uuid string, params param.UpdateAppBuildSystemParam) (*view.AppBuildSystemInventoryView, error) {
	var resp view.UpdateAppBuildSystemEventView
	if err := cli.Put("v1/appcenter/buildsystem/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAppBuildSystem deletes AppBuildSystem
func (cli *ZSClient) DeleteAppBuildSystem(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/buildsystem", uuid, string(deleteMode))
}
