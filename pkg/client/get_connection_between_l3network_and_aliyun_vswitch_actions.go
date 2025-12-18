// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetConnectionBetweenL3NetworkAndAliyunVSwitch gets ConnectionBetweenL3NetworkAndAliyunVSwitch by uuid
func (cli *ZSClient) GetConnectionBetweenL3NetworkAndAliyunVSwitch(uuid string) (*view.GetConnectionBetweenL3NetworkAndAliyunVSwitchView, error) {
	var resp view.GetConnectionBetweenL3NetworkAndAliyunVSwitchView
	if err := cli.Get("v1/hybrid/aliyun/relationships", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
