// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunProxyVSwitch updates AliyunProxyVSwitch
func (cli *ZSClient) UpdateAliyunProxyVSwitch(uuid string, params param.UpdateAliyunProxyVSwitchParam) (*view.AliyunProxyVSwitchInventoryView, error) {
	var resp view.UpdateAliyunProxyVSwitchEventView
	if err := cli.Put("v1/aliyun-proxy/vswitches/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAliyunProxyVSwitch queries AliyunProxyVSwitch list
func (cli *ZSClient) QueryAliyunProxyVSwitch(params *param.QueryParam) ([]view.AliyunProxyVSwitchInventoryView, error) {
	var resp []view.AliyunProxyVSwitchInventoryView
	return resp, cli.List("v1/aliyun-proxy/vpcs/vswitches", params, &resp)
}
// CreateAliyunProxyVSwitch creates AliyunProxyVSwitch
func (cli *ZSClient) CreateAliyunProxyVSwitch(params param.CreateAliyunProxyVSwitchParam) (*view.AliyunProxyVSwitchInventoryView, error) {
	var resp view.CreateAliyunProxyVSwitchEventView
	if err := cli.Post("v1/aliyun-proxy/vpcs/vswitches", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAliyunProxyVSwitch deletes AliyunProxyVSwitch
func (cli *ZSClient) DeleteAliyunProxyVSwitch(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/aliyun-proxy/vpcs/vswitches/{uuid}", uuid, string(deleteMode))
}
