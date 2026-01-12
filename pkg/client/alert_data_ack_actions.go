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
	var resp view.UpdateAlertDataAckEventView
	err := cli.PutWithSpec("v1/zwatch/alert-histories/acknowledgments", fmt.Sprintf(\"%s/actions\", alertDataUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
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
