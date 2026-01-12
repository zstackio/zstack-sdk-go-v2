// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAlarmLabel updates AlarmLabel
func (cli *ZSClient) UpdateAlarmLabel(uuid string, params param.UpdateAlarmLabelParam) (*view.AlarmLabelInventoryView, error) {
	var resp view.UpdateAlarmLabelEventView
	err := cli.PutWithSpec("v1/zwatch/alarms/labels", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
