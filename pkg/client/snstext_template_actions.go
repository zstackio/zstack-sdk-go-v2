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
	resp := view.SNSTextTemplateInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/alarms/sns/text-templates", uuid, "", map[string]interface{}{
		"updateSNSTextTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSTextTemplate creates SNSTextTemplate
func (cli *ZSClient) CreateSNSTextTemplate(params param.CreateSNSTextTemplateParam) (*view.SNSTextTemplateInventoryView, error) {
	resp := view.SNSTextTemplateInventoryView{}
	if err := cli.Post("v1/zwatch/alarms/sns/text-templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageSNSTextTemplate Pagination
func (cli *ZSClient) PageSNSTextTemplate(params *param.QueryParam) ([]view.SNSTextTemplateInventoryView, int, error) {
	var sNSTextTemplates []view.SNSTextTemplateInventoryView
	total, err := cli.Page("v1/zwatch/alarms/sns/text-templates", params, &sNSTextTemplates)
	return sNSTextTemplates, total, err
}
