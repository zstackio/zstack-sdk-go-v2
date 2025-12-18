// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetIAM2VirtualIDInGroup 获取IAM2VirtualIDInGroup详情
func (cli *ZSClient) GetIAM2VirtualIDInGroup(uuid string) (*view.GetIAM2VirtualIDInGroupView, error) {
	var resp view.GetIAM2VirtualIDInGroupView
	if err := cli.Get("v1/iam2/IAM2VirtualIDGroup/IAM2VirtualID", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

