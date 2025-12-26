// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeSchedulerState changes SchedulerState
func (cli *ZSClient) ChangeSchedulerState(uuid string, params param.ChangeSchedulerStateParam) (*view.ChangeSchedulerStateEventView, error) {
	resp := view.ChangeSchedulerStateEventView{}
	if err := cli.Put("v1/schedulers/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
