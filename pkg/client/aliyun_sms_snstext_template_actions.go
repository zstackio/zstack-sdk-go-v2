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
	resp := view.AliyunSmsSNSTextTemplateInventoryView{}
	if err := cli.Put("v1/zwatch/alarms/sns/text-templates", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateAliyunSmsSNSTextTemplate creates AliyunSmsSNSTextTemplate
func (cli *ZSClient) CreateAliyunSmsSNSTextTemplate(params param.CreateAliyunSmsSNSTextTemplateParam) (*view.SNSTextTemplateInventoryView, error) {
	resp := view.SNSTextTemplateInventoryView{}
	if err := cli.Post("v1/zwatch/alarms/sns/text-templates/aliyun-sms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAliyunSmsSNSTextTemplate queries AliyunSmsSNSTextTemplate list
func (cli *ZSClient) QueryAliyunSmsSNSTextTemplate(params *param.QueryParam) ([]view.AliyunSmsSNSTextTemplateInventoryView, error) {
	var resp []view.AliyunSmsSNSTextTemplateInventoryView
	return resp, cli.List("v1/zwatch/alarms/sns/text-templates/aliyun-sms", params, &resp)
}

// PageAliyunSmsSNSTextTemplate Pagination
func (cli *ZSClient) PageAliyunSmsSNSTextTemplate(params *param.QueryParam) ([]view.AliyunSmsSNSTextTemplateInventoryView, int, error) {
	var aliyunSmsSNSTextTemplates []view.AliyunSmsSNSTextTemplateInventoryView
	total, err := cli.Page("v1/zwatch/alarms/sns/text-templates/aliyun-sms", params, &aliyunSmsSNSTextTemplates)
	return aliyunSmsSNSTextTemplates, total, err
}
