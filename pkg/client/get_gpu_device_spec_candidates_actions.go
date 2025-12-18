// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetGpuDeviceSpecCandidates gets GpuDeviceSpecCandidates by uuid
func (cli *ZSClient) GetGpuDeviceSpecCandidates(uuid string) (*view.GetGpuDeviceSpecCandidatesView, error) {
	var resp view.GetGpuDeviceSpecCandidatesView
	if err := cli.Get("v1/gpu-device-specs/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
