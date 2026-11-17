// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// UpdateEmailMonitorTriggerAction updates EmailMonitorTrigger
func (cli *ZSClient) UpdateEmailMonitorTriggerAction(uuid string, params param.UpdateEmailMonitorTriggerActionParam) (*view.MonitorTriggerActionInventoryView, error) {
	resp := view.MonitorTriggerActionInventoryView{}
	if err := cli.PutWithRespKey("v1/monitoring/trigger-actions/emails", uuid, "", map[string]interface{}{
		"updateEmailMonitorTriggerAction": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEmailTriggerAction queries EmailTrigger list
func (cli *ZSClient) QueryEmailTriggerAction(params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, error) {
	var resp []view.MonitorTriggerActionInventoryView
	return resp, cli.List("v1/monitoring/trigger-actions/emails", params, &resp)
}

func (cli *ZSClient) GetEmailTriggerAction(uuid string) (*view.MonitorTriggerActionInventoryView, error) {
	var resp view.MonitorTriggerActionInventoryView
	if err := cli.Get("v1/monitoring/trigger-actions/emails", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEmailTriggerAction Pagination
func (cli *ZSClient) PageEmailTriggerAction(params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, int, error) {
	var emailTriggers []view.MonitorTriggerActionInventoryView
	total, err := cli.Page("v1/monitoring/trigger-actions/emails", params, &emailTriggers)
	return emailTriggers, total, err
}

// CreateEmailMonitorTriggerAction creates EmailMonitorTrigger
func (cli *ZSClient) CreateEmailMonitorTriggerAction(params param.CreateEmailMonitorTriggerActionParam) (*view.MonitorTriggerActionInventoryView, error) {
	resp := view.MonitorTriggerActionInventoryView{}
	if err := cli.Post("v1/monitoring/trigger-actions/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeMonitorTriggerActionState changes MonitorTriggerActionState
func (cli *ZSClient) ChangeMonitorTriggerActionState(uuid string, params param.ChangeMonitorTriggerActionStateParam) (*view.MonitorTriggerActionInventoryView, error) {
	resp := view.MonitorTriggerActionInventoryView{}
	if err := cli.PutWithRespKey("v1/monitoring/trigger-actions", uuid, "", map[string]interface{}{
		"changeMonitorTriggerActionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
