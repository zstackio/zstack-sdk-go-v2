// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSEndpointThirdpartyAlertHistory queries SNSEndpointThirdpartyAlertHistory list
func (cli *ZSClient) QuerySNSEndpointThirdpartyAlertHistory(params *param.QueryParam) ([]view.SNSEndpointThirdpartyAlertHistoryInventoryView, error) {
	var resp []view.SNSEndpointThirdpartyAlertHistoryInventoryView
	return resp, cli.List("v1/zwatch/third-party/alert-publish-histories", params, &resp)
}

func (cli *ZSClient) GetSNSEndpointThirdpartyAlertHistory(uuid string) (*view.SNSEndpointThirdpartyAlertHistoryInventoryView, error) {
	var resp view.SNSEndpointThirdpartyAlertHistoryInventoryView
	if err := cli.Get("v1/zwatch/third-party/alert-publish-histories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSEndpointThirdpartyAlertHistory Pagination
func (cli *ZSClient) PageSNSEndpointThirdpartyAlertHistory(params *param.QueryParam) ([]view.SNSEndpointThirdpartyAlertHistoryInventoryView, int, error) {
	var sNSEndpointThirdpartyAlertHistories []view.SNSEndpointThirdpartyAlertHistoryInventoryView
	total, err := cli.Page("v1/zwatch/third-party/alert-publish-histories", params, &sNSEndpointThirdpartyAlertHistories)
	return sNSEndpointThirdpartyAlertHistories, total, err
}

// QueryThirdpartyAlert queries ThirdpartyAlert list
func (cli *ZSClient) QueryThirdpartyAlert(params *param.QueryParam) ([]view.ThirdpartyOriginalAlertInventoryView, error) {
	var resp []view.ThirdpartyOriginalAlertInventoryView
	return resp, cli.List("v1/zwatch/third-party/alerts", params, &resp)
}

func (cli *ZSClient) GetThirdpartyAlert(uuid string) (*view.ThirdpartyOriginalAlertInventoryView, error) {
	var resp view.ThirdpartyOriginalAlertInventoryView
	if err := cli.Get("v1/zwatch/third-party/alerts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageThirdpartyAlert Pagination
func (cli *ZSClient) PageThirdpartyAlert(params *param.QueryParam) ([]view.ThirdpartyOriginalAlertInventoryView, int, error) {
	var thirdpartyAlerts []view.ThirdpartyOriginalAlertInventoryView
	total, err := cli.Page("v1/zwatch/third-party/alerts", params, &thirdpartyAlerts)
	return thirdpartyAlerts, total, err
}

// UpdateThirdpartyAlerts updates ThirdpartyAlerts
func (cli *ZSClient) UpdateThirdpartyAlerts(params param.UpdateThirdpartyAlertsParam) (*view.UpdateThirdpartyAlertsEventView, error) {
	resp := view.UpdateThirdpartyAlertsEventView{}
	if err := cli.PutWithRespKey("v1/zwatch/third-party/alerts/actions", "", "", map[string]interface{}{
		"updateThirdpartyAlerts": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
