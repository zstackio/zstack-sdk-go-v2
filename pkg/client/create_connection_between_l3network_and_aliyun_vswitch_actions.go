// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateConnectionBetweenL3NetworkAndAliyunVSwitch creates ConnectionBetweenL3NetworkAndAliyunVSwitch
func (cli *ZSClient) CreateConnectionBetweenL3NetworkAndAliyunVSwitch(params param.CreateConnectionBetweenL3NetworkAndAliyunVSwitchParam) (*view.CreateConnectionBetweenL3NetworkAndAliyunVSwitchEventView, error) {
	resp := view.CreateConnectionBetweenL3NetworkAndAliyunVSwitchEventView{}
	if err := cli.Post("v1/hybrid/aliyun/connections", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
