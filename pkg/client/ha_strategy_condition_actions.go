// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHaStrategyCondition updates HaStrategyCondition
func (cli *ZSClient) UpdateHaStrategyCondition(uuid string, params param.UpdateHaStrategyConditionParam) (*view.HaStrategyConditionInventoryView, error) {
	resp := view.HaStrategyConditionInventoryView{}
	if err := cli.PutWithRespKey("v1/ha-strategy-condition", uuid, "", map[string]interface{}{
		"updateHaStrategyCondition": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
