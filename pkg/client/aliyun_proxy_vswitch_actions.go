// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunProxyVSwitch updates AliyunProxyVSwitch
func (cli *ZSClient) UpdateAliyunProxyVSwitch(uuid string, params param.UpdateAliyunProxyVSwitchParam) (*view.AliyunProxyVSwitchInventoryView, error) {
	resp := view.AliyunProxyVSwitchInventoryView{}
	if err := cli.PutWithRespKey("v1/aliyun-proxy/vswitches", uuid, "", map[string]interface{}{
		"updateAliyunProxyVSwitch": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAliyunProxyVSwitch queries AliyunProxyVSwitch list
func (cli *ZSClient) QueryAliyunProxyVSwitch(params *param.QueryParam) ([]view.AliyunProxyVSwitchInventoryView, error) {
	var resp []view.AliyunProxyVSwitchInventoryView
	return resp, cli.List("v1/aliyun-proxy/vpcs/vswitches", params, &resp)
}

func (cli *ZSClient) GetAliyunProxyVSwitch(uuid string) (*view.AliyunProxyVSwitchInventoryView, error) {
	var resp view.AliyunProxyVSwitchInventoryView
	if err := cli.Get("v1/aliyun-proxy/vpcs/vswitches", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunProxyVSwitch Pagination
func (cli *ZSClient) PageAliyunProxyVSwitch(params *param.QueryParam) ([]view.AliyunProxyVSwitchInventoryView, int, error) {
	var aliyunProxyVSwitchs []view.AliyunProxyVSwitchInventoryView
	total, err := cli.Page("v1/aliyun-proxy/vpcs/vswitches", params, &aliyunProxyVSwitchs)
	return aliyunProxyVSwitchs, total, err
}
// CreateAliyunProxyVSwitch creates AliyunProxyVSwitch
func (cli *ZSClient) CreateAliyunProxyVSwitch(params param.CreateAliyunProxyVSwitchParam) (*view.AliyunProxyVSwitchInventoryView, error) {
	resp := view.AliyunProxyVSwitchInventoryView{}
	if err := cli.Post("v1/aliyun-proxy/vpcs/vswitches", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAliyunProxyVSwitch deletes AliyunProxyVSwitch
func (cli *ZSClient) DeleteAliyunProxyVSwitch(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/aliyun-proxy/vpcs/vswitches", uuid, string(deleteMode))
}
