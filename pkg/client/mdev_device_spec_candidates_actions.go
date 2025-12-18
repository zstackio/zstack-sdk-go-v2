// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetMdevDeviceSpecCandidates 获取MdevDeviceSpecCandidates详情
func (cli *ZSClient) GetMdevDeviceSpecCandidates(uuid string) (*view.GetMdevDeviceSpecCandidatesView, error) {
	var resp view.GetMdevDeviceSpecCandidatesView
	if err := cli.Get("v1/mdev-device-specs/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

