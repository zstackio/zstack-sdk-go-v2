// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAlarmLabel updates AlarmLabel
func (cli *ZSClient) UpdateAlarmLabel(uuid string, params param.UpdateAlarmLabelParam) (*view.AlarmLabelInventoryView, error) {
	var resp view.UpdateAlarmLabelEventView
	if err := cli.Put("v1/zwatch/alarms/labels/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
