// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSSnmpPlatform queries SNSSnmpPlatform list
func (cli *ZSClient) QuerySNSSnmpPlatform(params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, error) {
	var resp []view.SNSEmailPlatformInventoryView
	return resp, cli.List("v1/sns/application-platforms/snmp", params, &resp)
}
// UpdateSNSSnmpPlatform updates SNSSnmpPlatform
func (cli *ZSClient) UpdateSNSSnmpPlatform(uuid string, params param.UpdateSNSSnmpPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	var resp view.UpdateSNSApplicationPlatformEventView
	if err := cli.Put("v1/sns/application-platforms/snmp/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateSNSSnmpPlatform creates SNSSnmpPlatform
func (cli *ZSClient) CreateSNSSnmpPlatform(params param.CreateSNSSnmpPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	var resp view.CreateSNSApplicationPlatformEventView
	if err := cli.Post("v1/sns/application-platforms/snmp", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
