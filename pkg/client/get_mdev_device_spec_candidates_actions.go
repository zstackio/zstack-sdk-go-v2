// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetMdevDeviceSpecCandidates gets MdevDeviceSpecCandidates by uuid
func (cli *ZSClient) GetMdevDeviceSpecCandidates(uuid string) (*view.GetMdevDeviceSpecCandidatesView, error) {
	var resp view.GetMdevDeviceSpecCandidatesView
	if err := cli.Get("v1/mdev-device-specs/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
