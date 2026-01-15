// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAlertDataAck updates AlertDataAck
func (cli *ZSClient) UpdateAlertDataAck(alertDataUuid string, params param.UpdateAlertDataAckParam) (*view.AlertDataAckInventoryView, error) {
	resp := view.AlertDataAckInventoryView{}
	if err := cli.Put("v1/zwatch/alert-histories/acknowledgments", alertDataUuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAlertDataAck queries AlertDataAck list
func (cli *ZSClient) QueryAlertDataAck(params *param.QueryParam) ([]view.AlertDataAckInventoryView, error) {
	var resp []view.AlertDataAckInventoryView
	return resp, cli.List("v1/zwatch/alert-histories/acknowledgments", params, &resp)
}

func (cli *ZSClient) GetAlertDataAck(uuid string) (*view.AlertDataAckInventoryView, error) {
	var resp view.AlertDataAckInventoryView
	if err := cli.Get("v1/zwatch/alert-histories/acknowledgments", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAlertDataAck Pagination
func (cli *ZSClient) PageAlertDataAck(params *param.QueryParam) ([]view.AlertDataAckInventoryView, int, error) {
	var alertDataAcks []view.AlertDataAckInventoryView
	total, err := cli.Page("v1/zwatch/alert-histories/acknowledgments", params, &alertDataAcks)
	return alertDataAcks, total, err
}
