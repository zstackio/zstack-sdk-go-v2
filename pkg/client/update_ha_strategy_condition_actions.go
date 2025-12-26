// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateHaStrategyCondition updates HaStrategyCondition
func (cli *ZSClient) UpdateHaStrategyCondition(uuid string, params param.UpdateHaStrategyConditionParam) (*view.UpdateHaStrategyConditionEventView, error) {
	resp := view.UpdateHaStrategyConditionEventView{}
	if err := cli.Put("v1/ha-strategy-condition/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
