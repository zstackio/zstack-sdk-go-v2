// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeSchedulerState 操作SchedulerState
func (cli *ZSClient) ChangeSchedulerState(uuid string, params param.ChangeSchedulerStateParam) (*view.ChangeSchedulerStateEventView, error) {
	resp := view.ChangeSchedulerStateEventView{}
	if err := cli.Put("v1/schedulers/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

