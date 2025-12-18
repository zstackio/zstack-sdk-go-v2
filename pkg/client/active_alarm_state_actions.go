// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeActiveAlarmState 操作ActiveAlarmState
func (cli *ZSClient) ChangeActiveAlarmState(params param.ChangeActiveAlarmStateParam) (*view.ChangeActiveAlarmStateEventView, error) {
	resp := view.ChangeActiveAlarmStateEventView{}
	if err := cli.Post("v1/zwatch/activealarms/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

