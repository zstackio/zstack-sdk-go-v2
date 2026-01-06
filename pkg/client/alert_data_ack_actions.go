// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAlertDataAck updates AlertDataAck
func (cli *ZSClient) UpdateAlertDataAck(uuid string, params param.UpdateAlertDataAckParam) (*view.AlertDataAckInventoryView, error) {
	var resp view.UpdateAlertDataAckEventView
	if err := cli.Put("v1/zwatch/alert-histories/acknowledgments/{alertDataUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAlertDataAck queries AlertDataAck list
func (cli *ZSClient) QueryAlertDataAck(params *param.QueryParam) ([]view.AlertDataAckInventoryView, error) {
	var resp []view.AlertDataAckInventoryView
	return resp, cli.List("v1/zwatch/alert-histories/acknowledgments", params, &resp)
}
