// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSNSApplicationPlatform deletes SNSApplicationPlatform
func (cli *ZSClient) DeleteSNSApplicationPlatform(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-platforms", uuid, string(deleteMode))
}
// UpdateSNSApplicationPlatform updates SNSApplicationPlatform
func (cli *ZSClient) UpdateSNSApplicationPlatform(uuid string, params param.UpdateSNSApplicationPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	resp := view.SNSApplicationPlatformInventoryView{}
	if err := cli.PutWithRespKey("v1/sns/application-platforms", uuid, "", map[string]interface{}{
		"updateSNSApplicationPlatform": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSApplicationPlatform queries SNSApplicationPlatform list
func (cli *ZSClient) QuerySNSApplicationPlatform(params *param.QueryParam) ([]view.SNSApplicationPlatformInventoryView, error) {
	var resp []view.SNSApplicationPlatformInventoryView
	return resp, cli.List("v1/sns/application-platforms", params, &resp)
}

func (cli *ZSClient) GetSNSApplicationPlatform(uuid string) (*view.SNSApplicationPlatformInventoryView, error) {
	var resp view.SNSApplicationPlatformInventoryView
	if err := cli.Get("v1/sns/application-platforms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSApplicationPlatform Pagination
func (cli *ZSClient) PageSNSApplicationPlatform(params *param.QueryParam) ([]view.SNSApplicationPlatformInventoryView, int, error) {
	var sNSApplicationPlatforms []view.SNSApplicationPlatformInventoryView
	total, err := cli.Page("v1/sns/application-platforms", params, &sNSApplicationPlatforms)
	return sNSApplicationPlatforms, total, err
}
