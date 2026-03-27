// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAlarmLabel updates AlarmLabel
func (cli *ZSClient) UpdateAlarmLabel(ctx context.Context, uuid string, params param.UpdateAlarmLabelParam) (*view.AlarmLabelInventoryView, error) {
	resp := view.AlarmLabelInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/alarms/labels", uuid, "", map[string]interface{}{
		"updateAlarmLabel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
