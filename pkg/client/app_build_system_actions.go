// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAppBuildSystem adds AppBuildSystem
func (cli *ZSClient) AddAppBuildSystem(ctx context.Context, params param.AddAppBuildSystemParam) (*view.AppBuildSystemInventoryView, error) {
	resp := view.AppBuildSystemInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/appcenter/buildsystem", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAppBuildSystem queries AppBuildSystem list
func (cli *ZSClient) QueryAppBuildSystem(ctx context.Context, params *param.QueryParam) ([]view.AppBuildSystemInventoryView, error) {
	var resp []view.AppBuildSystemInventoryView
	return resp, cli.List(ctx, "v1/appcenter/buildsystem", params, &resp)
}

func (cli *ZSClient) GetAppBuildSystem(ctx context.Context, uuid string) (*view.AppBuildSystemInventoryView, error) {
	var resp view.AppBuildSystemInventoryView
	if err := cli.Get(ctx, "v1/appcenter/buildsystem", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAppBuildSystem Pagination
func (cli *ZSClient) PageAppBuildSystem(ctx context.Context, params *param.QueryParam) ([]view.AppBuildSystemInventoryView, int, error) {
	var appBuildSystems []view.AppBuildSystemInventoryView
	total, err := cli.Page(ctx, "v1/appcenter/buildsystem", params, &appBuildSystems)
	return appBuildSystems, total, err
}
// ReconnectAppBuildSystem operates on AppBuildSystem
func (cli *ZSClient) ReconnectAppBuildSystem(ctx context.Context, uuid string, params param.ReconnectAppBuildSystemParam) (*view.AppBuildSystemInventoryView, error) {
	resp := view.AppBuildSystemInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/appcenter/buildsystem", uuid, "", map[string]interface{}{
		"reconnectAppBuildSystem": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateAppBuildSystem updates AppBuildSystem
func (cli *ZSClient) UpdateAppBuildSystem(ctx context.Context, uuid string, params param.UpdateAppBuildSystemParam) (*view.AppBuildSystemInventoryView, error) {
	resp := view.AppBuildSystemInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/appcenter/buildsystem", uuid, "", map[string]interface{}{
		"updateAppBuildSystem": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAppBuildSystem deletes AppBuildSystem
func (cli *ZSClient) DeleteAppBuildSystem(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/appcenter/buildsystem", uuid, string(deleteMode))
}
