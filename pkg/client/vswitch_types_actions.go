// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVSwitchTypes 获取VSwitchTypes详情
func (cli *ZSClient) GetVSwitchTypes(uuid string) (*view.GetVSwitchTypesView, error) {
	var resp view.GetVSwitchTypesView
	if err := cli.Get("v1/l2-networks/vSwitchTypes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

