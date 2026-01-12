// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunSmsSNSTextTemplate updates AliyunSmsSNSTextTemplate
func (cli *ZSClient) UpdateAliyunSmsSNSTextTemplate(uuid string, params param.UpdateAliyunSmsSNSTextTemplateParam) (*view.AliyunSmsSNSTextTemplateInventoryView, error) {
	var resp view.UpdateAliyunSmsSNSTextTemplateEventView
	if err := cli.Put("v1/zwatch/alarms/sns/text-templates", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateAliyunSmsSNSTextTemplate creates AliyunSmsSNSTextTemplate
func (cli *ZSClient) CreateAliyunSmsSNSTextTemplate(params param.CreateAliyunSmsSNSTextTemplateParam) (*view.SNSTextTemplateInventoryView, error) {
	var resp view.CreateSNSTextTemplateEventView
	if err := cli.Post("v1/zwatch/alarms/sns/text-templates/aliyun-sms", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAliyunSmsSNSTextTemplate queries AliyunSmsSNSTextTemplate list
func (cli *ZSClient) QueryAliyunSmsSNSTextTemplate(params *param.QueryParam) ([]view.AliyunSmsSNSTextTemplateInventoryView, error) {
	var resp []view.AliyunSmsSNSTextTemplateInventoryView
	return resp, cli.List("v1/zwatch/alarms/sns/text-templates/aliyun-sms", params, &resp)
}

func (cli *ZSClient) GetAliyunSmsSNSTextTemplate(uuid string) (*view.AliyunSmsSNSTextTemplateInventoryView, error) {
	var resp view.AliyunSmsSNSTextTemplateInventoryView
	if err := cli.Get("v1/zwatch/alarms/sns/text-templates/aliyun-sms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
