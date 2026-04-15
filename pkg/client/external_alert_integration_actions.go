// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{}    // avoid unused import

// SetupExternalAlertIntegration 创建或更新外部告警集成的 HTTP Endpoint、订阅及推送模板（幂等）
func (cli *ZSClient) SetupExternalAlertIntegration(params param.SetupExternalAlertIntegrationParam) (*view.SetupExternalAlertIntegrationView, error) {
	resp := view.SetupExternalAlertIntegrationView{}
	if err := cli.Post("v1/sns/external-alert-integrations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckExternalAlertIntegration 检查外部告警集成健康状态
func (cli *ZSClient) CheckExternalAlertIntegration(params param.CheckExternalAlertIntegrationParam) (*view.CheckExternalAlertIntegrationView, error) {
	resp := view.CheckExternalAlertIntegrationView{}
	if err := cli.Post("v1/sns/external-alert-integrations/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveExternalAlertIntegration 移除外部告警集成（删除 Endpoint、订阅及推送模板）
func (cli *ZSClient) RemoveExternalAlertIntegration(params param.RemoveExternalAlertIntegrationParam) (*view.RemoveExternalAlertIntegrationView, error) {
	resp := view.RemoveExternalAlertIntegrationView{}
	if err := cli.Post("v1/sns/external-alert-integrations/remove", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
