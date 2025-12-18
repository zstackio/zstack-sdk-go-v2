// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch updates ConnectionBetweenL3NetWorkAndAliyunVSwitch
func (cli *ZSClient) UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch(uuid string, params param.UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParam) (*view.UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchEventView, error) {
	resp := view.UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchEventView{}
	if err := cli.Put("v1/hybrid/aliyun/connections/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
