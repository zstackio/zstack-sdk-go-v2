// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetMdevDeviceCandidates 获取MdevDeviceCandidates详情
func (cli *ZSClient) GetMdevDeviceCandidates(uuid string) (*view.GetMdevDeviceCandidatesView, error) {
	var resp view.GetMdevDeviceCandidatesView
	if err := cli.Get("v1/mdev-devices/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

