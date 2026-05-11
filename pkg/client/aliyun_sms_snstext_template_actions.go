// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunSmsSNSTextTemplate updates AliyunSmsSNSTextTemplate
func (cli *ZSClient) UpdateAliyunSmsSNSTextTemplate(ctx context.Context, uuid string, params param.UpdateAliyunSmsSNSTextTemplateParam) (*view.AliyunSmsSNSTextTemplateInventoryView, error) {
	resp := view.AliyunSmsSNSTextTemplateInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/alarms/sns/text-templates", uuid, "", map[string]interface{}{
		"updateAliyunSmsSNSTextTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateAliyunSmsSNSTextTemplate creates AliyunSmsSNSTextTemplate
func (cli *ZSClient) CreateAliyunSmsSNSTextTemplate(ctx context.Context, params param.CreateAliyunSmsSNSTextTemplateParam) (*view.SNSTextTemplateInventoryView, error) {
	resp := view.SNSTextTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/zwatch/alarms/sns/text-templates/aliyun-sms", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAliyunSmsSNSTextTemplate queries AliyunSmsSNSTextTemplate list
func (cli *ZSClient) QueryAliyunSmsSNSTextTemplate(ctx context.Context, params *param.QueryParam) ([]view.AliyunSmsSNSTextTemplateInventoryView, error) {
	var resp []view.AliyunSmsSNSTextTemplateInventoryView
	return resp, cli.List(ctx, "v1/zwatch/alarms/sns/text-templates/aliyun-sms", params, &resp)
}

func (cli *ZSClient) GetAliyunSmsSNSTextTemplate(ctx context.Context, uuid string) (*view.AliyunSmsSNSTextTemplateInventoryView, error) {
	var resp view.AliyunSmsSNSTextTemplateInventoryView
	if err := cli.Get(ctx, "v1/zwatch/alarms/sns/text-templates/aliyun-sms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunSmsSNSTextTemplate Pagination
func (cli *ZSClient) PageAliyunSmsSNSTextTemplate(ctx context.Context, params *param.QueryParam) ([]view.AliyunSmsSNSTextTemplateInventoryView, int, error) {
	var aliyunSmsSNSTextTemplates []view.AliyunSmsSNSTextTemplateInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/alarms/sns/text-templates/aliyun-sms", params, &aliyunSmsSNSTextTemplates)
	return aliyunSmsSNSTextTemplates, total, err
}
