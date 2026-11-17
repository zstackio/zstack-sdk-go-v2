// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryManagementNode queries ManagementNode list
func (cli *ZSClient) QueryManagementNode(params *param.QueryParam) ([]view.ManagementNodeInventoryView, error) {
	var resp []view.ManagementNodeInventoryView
	return resp, cli.List("v1/management-nodes", params, &resp)
}

func (cli *ZSClient) GetManagementNode(uuid string) (*view.ManagementNodeInventoryView, error) {
	var resp view.ManagementNodeInventoryView
	if err := cli.Get("v1/management-nodes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageManagementNode Pagination
func (cli *ZSClient) PageManagementNode(params *param.QueryParam) ([]view.ManagementNodeInventoryView, int, error) {
	var managementNodes []view.ManagementNodeInventoryView
	total, err := cli.Page("v1/management-nodes", params, &managementNodes)
	return managementNodes, total, err
}

// GetVersion gets Version by uuid
func (cli *ZSClient) GetVersion() (*view.GetVersionView, error) {
	resp := view.GetVersionView{}
	if err := cli.PutWithRespKey("v1/management-nodes/actions", "", "", map[string]interface{}{
		"getVersion": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetManagementNodeDirCapacity gets ManagementNodeDirCapacity by uuid
func (cli *ZSClient) GetManagementNodeDirCapacity() (*view.GetManagementNodeDirCapacityView, error) {
	var resp view.GetManagementNodeDirCapacityView
	if err := cli.GetWithRespKey("v1/zwatch/mn", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFactoryModeState gets FactoryModeState by uuid
func (cli *ZSClient) GetFactoryModeState() (*view.GetFactoryModeStateView, error) {
	var resp view.GetFactoryModeStateView
	if err := cli.GetWithRespKey("v1/management-nodes/factory-mode-state", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFactoryModeState updates FactoryModeState
func (cli *ZSClient) UpdateFactoryModeState(params param.UpdateFactoryModeStateParam) (*view.UpdateFactoryModeStateEventView, error) {
	resp := view.UpdateFactoryModeStateEventView{}
	if err := cli.PutWithRespKey("v1/management-nodes/actions", "", "", map[string]interface{}{
		"updateFactoryModeState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSupportAPIs gets SupportAPIs by uuid
func (cli *ZSClient) GetSupportAPIs() (*view.GetSupportAPIsView, error) {
	resp := view.GetSupportAPIsView{}
	if err := cli.PutWithRespKey("v1/management-nodes/actions", "", "", map[string]interface{}{
		"getSupportAPIs": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsReadyToGo operates on IsReadyToGo
func (cli *ZSClient) IsReadyToGo() (*view.IsReadyToGoView, error) {
	var resp view.IsReadyToGoView
	if err := cli.GetWithRespKey("v1/management-nodes/ready", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetManagementNodeArch gets ManagementNodeArch by uuid
func (cli *ZSClient) GetManagementNodeArch() (*view.GetManagementNodeArchView, error) {
	resp := view.GetManagementNodeArchView{}
	if err := cli.PutWithRespKey("v1/management-nodes/actions", "", "", map[string]interface{}{
		"getManagementNodeArch": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetManagementNodeOS gets ManagementNodeOS by uuid
func (cli *ZSClient) GetManagementNodeOS() (*view.GetManagementNodeOSView, error) {
	resp := view.GetManagementNodeOSView{}
	if err := cli.PutWithRespKey("v1/management/actions", "", "", map[string]interface{}{
		"getManagementNodeOS": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPlatformTimeZone gets PlatformTimeZone by uuid
func (cli *ZSClient) GetPlatformTimeZone() (*view.GetPlatformTimeZoneView, error) {
	var resp view.GetPlatformTimeZoneView
	if err := cli.GetWithRespKey("v1/management-nodes/platform-timezone", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsOpensourceVersion operates on IsOpensourceVersion
func (cli *ZSClient) IsOpensourceVersion() (*view.IsOpensourceVersionView, error) {
	var resp view.IsOpensourceVersionView
	if err := cli.GetWithRespKey("v1/meta-data/opensource", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCurrentTime gets CurrentTime by uuid
func (cli *ZSClient) GetCurrentTime() (*view.GetCurrentTimeView, error) {
	resp := view.GetCurrentTimeView{}
	if err := cli.PutWithRespKey("v1/management-nodes/actions", "", "", map[string]interface{}{
		"getCurrentTime": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
