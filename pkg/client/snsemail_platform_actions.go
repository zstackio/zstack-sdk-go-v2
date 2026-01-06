// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSEmailPlatform creates SNSEmailPlatform
func (cli *ZSClient) CreateSNSEmailPlatform(params param.CreateSNSEmailPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	var resp view.CreateSNSApplicationPlatformEventView
	if err := cli.Post("v1/sns/application-platforms/email", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ValidateSNSEmailPlatform operates on SNSEmailPlatform
func (cli *ZSClient) ValidateSNSEmailPlatform(uuid string, params param.ValidateSNSEmailPlatformParam) (*view.SNSEmailPlatformInventoryView, error) {
	resp := view.SNSEmailPlatformInventoryView{}
	if err := cli.Put("v1/sns/application-platforms/email/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSEmailPlatform queries SNSEmailPlatform list
func (cli *ZSClient) QuerySNSEmailPlatform(params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, error) {
	var resp []view.SNSEmailPlatformInventoryView
	return resp, cli.List("v1/sns/application-platforms/email", params, &resp)
}
