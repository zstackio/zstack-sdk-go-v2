// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAppBuildSystem adds AppBuildSystem
func (cli *ZSClient) AddAppBuildSystem(params param.AddAppBuildSystemParam) (*view.AppBuildSystemInventoryView, error) {
	resp := view.AppBuildSystemInventoryView{}
	if err := cli.Post("v1/appcenter/buildsystem", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageAppBuildSystem Pagination
func (cli *ZSClient) PageAppBuildSystem(params *param.QueryParam) ([]view.AppBuildSystemInventoryView, int, error) {
	var appBuildSystems []view.AppBuildSystemInventoryView
	total, err := cli.Page("v1/appcenter/buildsystem", params, &appBuildSystems)
	return appBuildSystems, total, err
}
// ReconnectAppBuildSystem operates on AppBuildSystem
func (cli *ZSClient) ReconnectAppBuildSystem(uuid string, params param.ReconnectAppBuildSystemParam) (*view.AppBuildSystemInventoryView, error) {
	resp := view.AppBuildSystemInventoryView{}
	if err := cli.PutWithRespKey("v1/appcenter/buildsystem", uuid, "", map[string]interface{}{
		"reconnectAppBuildSystem": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateAppBuildSystem updates AppBuildSystem
func (cli *ZSClient) UpdateAppBuildSystem(uuid string, params param.UpdateAppBuildSystemParam) (*view.AppBuildSystemInventoryView, error) {
	resp := view.AppBuildSystemInventoryView{}
	if err := cli.PutWithRespKey("v1/appcenter/buildsystem", uuid, "", map[string]interface{}{
		"updateAppBuildSystem": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAppBuildSystem deletes AppBuildSystem
func (cli *ZSClient) DeleteAppBuildSystem(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/buildsystem", uuid, string(deleteMode))
}

// ChangeAppBuildSystemState changes AppBuildSystemState
func (cli *ZSClient) ChangeAppBuildSystemState(uuid string, params param.ChangeAppBuildSystemStateParam) (*view.AppBuildSystemInventoryView, error) {
	resp := view.AppBuildSystemInventoryView{}
	if err := cli.PutWithRespKey("v1/appcenter/buildsystem", uuid, "", map[string]interface{}{
		"changeAppBuildSystemState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachAppBuildSystemToZone operates on AppBuildSystemToZone
func (cli *ZSClient) AttachAppBuildSystemToZone(zoneUuid string, buildSystemUuid string, params param.AttachAppBuildSystemToZoneParam) (*view.AppBuildSystemZoneRefInventoryView, error) {
	resp := view.AppBuildSystemZoneRefInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/zones/%s/buildsystem/%s", zoneUuid, buildSystemUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachAppBuildSystemToZone operates on AppBuildSystemToZone
func (cli *ZSClient) DetachAppBuildSystemToZone(zoneUuid string, buildSystemUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/zones", zoneUuid, fmt.Sprintf("buildsystem/%s", buildSystemUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetAppBuildSystemCapacity gets AppBuildSystemCapacity by uuid
func (cli *ZSClient) GetAppBuildSystemCapacity() (*view.GetAppBuildSystemCapacityView, error) {
	var resp view.GetAppBuildSystemCapacityView
	if err := cli.GetWithRespKey("v1/appcenter/buildsystem/capacities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
