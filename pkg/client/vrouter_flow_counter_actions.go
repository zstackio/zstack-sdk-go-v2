// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVRouterFlowCounter 获取VRouterFlowCounter详情
func (cli *ZSClient) GetVRouterFlowCounter(uuid string) (*view.GetVRouterFlowCounterView, error) {
	var resp view.GetVRouterFlowCounterView
	if err := cli.Get("v1/flowmeters/{vRouterUuid}/counter", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

