// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAlarmLabel updates AlarmLabel
func (cli *ZSClient) UpdateAlarmLabel(uuid string, params param.UpdateAlarmLabelParam) (*view.AlarmLabelInventoryView, error) {
	resp := view.AlarmLabelInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/alarms/labels", uuid, "", map[string]interface{}{
		"updateAlarmLabel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddLabelToAlarm adds LabelToAlarm
func (cli *ZSClient) AddLabelToAlarm(alarmUuid string, params param.AddLabelToAlarmParam) (*view.AlarmLabelInventoryView, error) {
	resp := view.AlarmLabelInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/zwatch/alarms/%s/labels", alarmUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveLabelFromAlarm removes LabelFromAlarm
func (cli *ZSClient) RemoveLabelFromAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/labels", uuid, string(deleteMode))
}
