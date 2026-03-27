// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSSnmpPlatform queries SNSSnmpPlatform list
func (cli *ZSClient) QuerySNSSnmpPlatform(ctx context.Context, params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, error) {
	var resp []view.SNSEmailPlatformInventoryView
	return resp, cli.List(ctx, "v1/sns/application-platforms/snmp", params, &resp)
}

func (cli *ZSClient) GetSNSSnmpPlatform(ctx context.Context, uuid string) (*view.SNSEmailPlatformInventoryView, error) {
	var resp view.SNSEmailPlatformInventoryView
	if err := cli.Get(ctx, "v1/sns/application-platforms/snmp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSSnmpPlatform Pagination
func (cli *ZSClient) PageSNSSnmpPlatform(ctx context.Context, params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, int, error) {
	var sNSSnmpPlatforms []view.SNSEmailPlatformInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-platforms/snmp", params, &sNSSnmpPlatforms)
	return sNSSnmpPlatforms, total, err
}
// UpdateSNSSnmpPlatform updates SNSSnmpPlatform
func (cli *ZSClient) UpdateSNSSnmpPlatform(ctx context.Context, uuid string, params param.UpdateSNSSnmpPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	resp := view.SNSApplicationPlatformInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-platforms/snmp", uuid, "", map[string]interface{}{
		"updateSNSSnmpPlatform": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSSnmpPlatform creates SNSSnmpPlatform
func (cli *ZSClient) CreateSNSSnmpPlatform(ctx context.Context, params param.CreateSNSSnmpPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	resp := view.SNSApplicationPlatformInventoryView{}
	if err := cli.Post(ctx, "v1/sns/application-platforms/snmp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
