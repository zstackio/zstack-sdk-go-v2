// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2Instance queries BareMetal2Instance list
func (cli *ZSClient) QueryBareMetal2Instance(params *param.QueryParam) ([]view.BareMetal2InstanceInventoryView, error) {
	var resp []view.BareMetal2InstanceInventoryView
	return resp, cli.List("v1/baremetal2/bm-instances", params, &resp)
}

func (cli *ZSClient) GetBareMetal2Instance(uuid string) (*view.BareMetal2InstanceInventoryView, error) {
	var resp view.BareMetal2InstanceInventoryView
	if err := cli.Get("v1/baremetal2/bm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2Instance Pagination
func (cli *ZSClient) PageBareMetal2Instance(params *param.QueryParam) ([]view.BareMetal2InstanceInventoryView, int, error) {
	var bareMetal2Instances []view.BareMetal2InstanceInventoryView
	total, err := cli.Page("v1/baremetal2/bm-instances", params, &bareMetal2Instances)
	return bareMetal2Instances, total, err
}
// CreateBareMetal2Instance creates BareMetal2Instance
func (cli *ZSClient) CreateBareMetal2Instance(params param.CreateBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.Post("v1/baremetal2/bm-instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBareMetal2InstanceAsync Async
func (cli *ZSClient) CreateBareMetal2InstanceAsync(params param.CreateBareMetal2InstanceParam) (string, error) {

	resource := "v1/baremetal2/bm-instances"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// ReconnectBareMetal2Instance operates on BareMetal2Instance
func (cli *ZSClient) ReconnectBareMetal2Instance(uuid string, params param.ReconnectBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/bm-instances", uuid, "", map[string]interface{}{
		"reconnectBareMetal2Instance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StartBareMetal2Instance starts BareMetal2Instance
func (cli *ZSClient) StartBareMetal2Instance(uuid string, params param.StartBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/bm-instances", uuid, "", map[string]interface{}{
		"startBareMetal2Instance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBareMetal2Instance updates BareMetal2Instance
func (cli *ZSClient) UpdateBareMetal2Instance(uuid string, params param.UpdateBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/bm-instances", uuid, "", map[string]interface{}{
		"updateBareMetal2Instance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBareMetal2SupportedBootMode gets BareMetal2SupportedBootMode by uuid
func (cli *ZSClient) GetBareMetal2SupportedBootMode() (*view.GetBareMetal2SupportedBootModeView, error) {
	var resp view.GetBareMetal2SupportedBootModeView
	if err := cli.GetWithRespKey("v1/baremetal2/chassis/supported-boot-modes", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2InstancePassword changes BareMetal2InstancePassword
func (cli *ZSClient) ChangeBareMetal2InstancePassword(uuid string, params param.ChangeBareMetal2InstancePasswordParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/bm-instances", uuid, "", map[string]interface{}{
		"changeBareMetal2InstancePassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckStaticProvisionIp operates on StaticProvisionIp
func (cli *ZSClient) CheckStaticProvisionIp(params param.CheckStaticProvisionIpParam) (*view.CheckStaticProvisionIpView, error) {
	resp := view.CheckStaticProvisionIpView{}
	if err := cli.Post("v1/baremetal2/bm-instances/static/provision/ip/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBareMetal2InstanceFromVolumeBackup creates BareMetal2InstanceFromVolumeBackup
func (cli *ZSClient) CreateBareMetal2InstanceFromVolumeBackup(backupUuid string, params param.CreateBareMetal2InstanceFromVolumeBackupParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/baremetal2-instance/from/volume-backup/%s", backupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBareMetal2InstanceFromVmBackup creates BareMetal2InstanceFromVmBackup
func (cli *ZSClient) CreateBareMetal2InstanceFromVmBackup(groupUuid string, params param.CreateBareMetal2InstanceFromVmBackupParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/baremetal2-instance/from/vm-backup/%s", groupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
