// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSNSApplicationPlatform deletes SNSApplicationPlatform
func (cli *ZSClient) DeleteSNSApplicationPlatform(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-platforms/{uuid}", uuid, string(deleteMode))
}
// UpdateSNSApplicationPlatform updates SNSApplicationPlatform
func (cli *ZSClient) UpdateSNSApplicationPlatform(uuid string, params param.UpdateSNSApplicationPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	var resp view.UpdateSNSApplicationPlatformEventView
	if err := cli.Put("v1/sns/application-platforms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySNSApplicationPlatform queries SNSApplicationPlatform list
func (cli *ZSClient) QuerySNSApplicationPlatform(params *param.QueryParam) ([]view.SNSApplicationPlatformInventoryView, error) {
	var resp []view.SNSApplicationPlatformInventoryView
	return resp, cli.List("v1/sns/application-platforms", params, &resp)
}
