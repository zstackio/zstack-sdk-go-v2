// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateSNSTextTemplate updates SNSTextTemplate
func (cli *ZSClient) UpdateSNSTextTemplate(uuid string, params param.UpdateSNSTextTemplateParam) (*view.SNSTextTemplateInventoryView, error) {
	var resp view.UpdateSNSTextTemplateEventView
	if err := cli.Put("v1/zwatch/alarms/sns/text-templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateSNSTextTemplate creates SNSTextTemplate
func (cli *ZSClient) CreateSNSTextTemplate(params param.CreateSNSTextTemplateParam) (*view.SNSTextTemplateInventoryView, error) {
	var resp view.CreateSNSTextTemplateEventView
	if err := cli.Post("v1/zwatch/alarms/sns/text-templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteSNSTextTemplate deletes SNSTextTemplate
func (cli *ZSClient) DeleteSNSTextTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/sns/text-templates", uuid, string(deleteMode))
}
// QuerySNSTextTemplate queries SNSTextTemplate list
func (cli *ZSClient) QuerySNSTextTemplate(params *param.QueryParam) ([]view.SNSTextTemplateInventoryView, error) {
	var resp []view.SNSTextTemplateInventoryView
	return resp, cli.List("v1/zwatch/alarms/sns/text-templates", params, &resp)
}

func (cli *ZSClient) GetSNSTextTemplate(uuid string) (*view.SNSTextTemplateInventoryView, error) {
	var resp view.SNSTextTemplateInventoryView
	if err := cli.Get("v1/zwatch/alarms/sns/text-templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
