// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSSnmpPlatform queries SNSSnmpPlatform list
func (cli *ZSClient) QuerySNSSnmpPlatform(params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, error) {
	var resp []view.SNSEmailPlatformInventoryView
	return resp, cli.List("v1/sns/application-platforms/snmp", params, &resp)
}

// PageSNSSnmpPlatform Pagination
func (cli *ZSClient) PageSNSSnmpPlatform(params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, int, error) {
	var sNSSnmpPlatforms []view.SNSEmailPlatformInventoryView
	total, err := cli.Page("v1/sns/application-platforms/snmp", params, &sNSSnmpPlatforms)
	return sNSSnmpPlatforms, total, err
}
// UpdateSNSSnmpPlatform updates SNSSnmpPlatform
func (cli *ZSClient) UpdateSNSSnmpPlatform(uuid string, params param.UpdateSNSSnmpPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	resp := view.SNSApplicationPlatformInventoryView{}
	if err := cli.Put("v1/sns/application-platforms/snmp", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSSnmpPlatform creates SNSSnmpPlatform
func (cli *ZSClient) CreateSNSSnmpPlatform(params param.CreateSNSSnmpPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	resp := view.SNSApplicationPlatformInventoryView{}
	if err := cli.Post("v1/sns/application-platforms/snmp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
